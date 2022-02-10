package mmGit

import (
	"GoSungro/Only"
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)


func (z *Git) Connect() error {

	for range Only.Once {
		if z.IsNotOk() {
			break
		}

		var ok bool
		ok, z.Error = IsDirExists(z.RepoDir)
		if z.Error != nil {
			break
		}
		if ok {
			z.Error = z.Open()
			break
		}

		z.Error = z.Clone()
		if z.Error != nil {
			break
		}
	}

	return z.Error
}

func (z *Git) Open() error {

	for range Only.Once {
		if z.IsNotOk() {
			break
		}

		z.repo, z.Error = git.PlainOpen(z.RepoDir)
		if z.Error != nil {
			break
		}

		z.worktree, z.Error = z.repo.Worktree()
		if z.Error != nil {
			break
		}

		var ref *plumbing.Reference
		ref, z.Error = z.repo.Head()
		if z.Error != nil {
			break
		}

		if ref.Hash().IsZero() {
			z.Error = errors.New("invalid HEAD reference")
			break
		}

		fmt.Printf("Git opened\n\trepo: %s\n\tdir: %s\n", z.RepoUrl, z.RepoDir)
	}

	return z.Error
}

func (z *Git) Clone() error {

	for range Only.Once {
		if z.IsNotOk() {
			break
		}

		var ok bool
		ok, z.Error = IsDirExists(z.RepoDir)
		if z.Error != nil {
			break
		}
		if ok {
			z.Error = errors.New(fmt.Sprintf("Cannot clone - directory '%s' already exists.", z.RepoDir))
			break
		}

		// CONTEXT: Provide Ctrl-C capability as well as operation timeouts.
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go func() {
			<-stop
			fmt.Println("\nCanceling operation...")
			cancel()
		}()
		// CONTEXT: Provide Ctrl-C capability as well as operation timeouts.

		pk := z.GetSshAuth()
		if z.Error != nil {
			break
		}

		options := &git.CloneOptions {
			URL:               z.RepoUrl,
			Auth:              pk,
			RemoteName:        "",
			ReferenceName:     "",
			SingleBranch:      false,
			NoCheckout:        false,
			Depth:             0,
			RecurseSubmodules: 0,
			Progress:          os.Stdout,
			Tags:              0,
			InsecureSkipTLS:   false,
			CABundle:          nil,
		}
		z.repo, z.Error = git.PlainCloneContext(ctx, z.RepoDir, false, options)
		if z.Error != nil {
			break
		}
		z.worktree, z.Error = z.repo.Worktree()
		if z.Error != nil {
			break
		}

		var ref *plumbing.Reference
		ref, z.Error = z.repo.Head()
		if z.Error != nil {
			break
		}

		if ref.Hash().IsZero() {
			z.Error = errors.New("invalid HEAD reference")
			break
		}

		fmt.Printf("Git cloned\n\trepo: %s\n\tdir: %s\n", z.RepoUrl, z.RepoDir)
	}

	return z.Error
}

// GetSshAuth: Gitlab keys need to be created with at least 3072 bits.
// ssh-keygen -t rsa -b 3072 -C 'root@everywhere' -f gitlab_rsa -N ''
func (z *Git) GetSshAuth() *ssh.PublicKeys {
	var pk *ssh.PublicKeys

	for range Only.Once {
		if z.IsNotOk() {
			break
		}

		var u *user.User
		u, z.Error = user.Current()

		paths := []string {
			z.KeyFile,
			filepath.Join(u.HomeDir, ".ssh", "id_rsa"),
		}

		var path string
		for _, path = range paths {
			if path == "" {
				continue
			}

			z.Error = checkKeyFile(path)
			if z.Error != nil {
				continue
			}

			break
		}

		// Try without password first.
		var password string
		pk, z.Error = ssh.NewPublicKeysFromFile("git", path, password)
		if z.Error == nil {
			fmt.Printf("AUTH: %v\n", pk)
			break
		}

		// Then with password.
		password = getPassword("ApiPassword: ")
		pk, z.Error = ssh.NewPublicKeysFromFile("git", path, password)
		if z.Error == nil {
			fmt.Printf("AUTH: %v\n", pk)
			break
		}
	}

	return pk
}

func checkKeyFile(path string) error {
	var err error

	for range Only.Once {
		if path == "" {
			continue
		}

		var fi os.FileInfo
		fi, err = os.Stat(path)
		if os.IsNotExist(err) {
			continue
		}

		if fi.IsDir() {
			err = errors.New("SSH publickey file is a directory")
			continue
		}
	}

	return err
}

// techEcho() - turns terminal echo on or off.
func termEcho(on bool) {
	// Common settings and variables for both stty calls.
	attrs := syscall.ProcAttr{
		Dir:   "",
		Env:   []string{},
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
		Sys:   nil}
	var ws syscall.WaitStatus
	cmd := "echo"
	if on == false {
		cmd = "-echo"
	}

	// Enable/disable echoing.
	pid, err := syscall.ForkExec(
		"/bin/stty",
		[]string{"stty", cmd},
		&attrs)
	if err != nil {
		panic(err)
	}

	// Wait for the stty process to complete.
	_, err = syscall.Wait4(pid, &ws, 0, nil)
	if err != nil {
		panic(err)
	}
}

// getPassword - Prompt for password.
func getPassword(prompt string) string {
	fmt.Print(prompt)

	// Catch a ^C interrupt.
	// Make sure that we reset term echo before exiting.
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		for _ = range signalChannel {
			fmt.Println("\n^C interrupt.")
			termEcho(true)
			os.Exit(1)
		}
	}()

	// Echo is disabled, now grab the data.
	termEcho(false) // disable terminal echo
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	termEcho(true) // always re-enable terminal echo
	fmt.Println("")
	if err != nil {
		// The terminal has been reset, go ahead and exit.
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}
	return strings.TrimSpace(text)
}

//func (z *Git) setContext() error {
//
//	for range Only.Once {
//		if z.IsNotOk() {
//			break
//		}
//
//		stop := make(chan os.Signal, 1)
//		signal.Notify(stop, os.Interrupt)
//		ctx, cancel := context.WithCancel(context.Background())
//		defer cancel() // cancel when we are finished consuming integers
//
//		go func() {
//			<-stop
//			Warning("\nSignal detected, canceling operation...")
//			cancel()
//		}()
//
//		var auth ssh.AuthMethod
//		auth, z.Error = ssh.DefaultAuthBuilder("admin-mickh")
//		if z.Error != nil {
//			break
//		}
//
//		z.repo, z.Error = git.PlainClone(z.RepoDir, false, &git.CloneOptions {
//			URL:  z.RepoUrl,
//			Auth: auth,
//		})
//		if z.Error != nil {
//			break
//		}
//
//		var ref *plumbing.Reference
//		ref, z.Error = z.repo.Head()
//		if z.Error != nil {
//			break
//		}
//
//		var commit *object.Commit
//		commit, z.Error = z.repo.CommitObject(ref.Hash())
//		if z.Error != nil {
//			break
//		}
//
//		fmt.Println(commit)
//	}
//
//	return z.Error
//}

func (z *Git) SaveFile(fn string, data []byte) error {

	for range Only.Once {
		if z.IsNotOk() {
			break
		}

		//z.worktree, z.Error = z.repo.Worktree()
		//if z.Error != nil {
		//	break
		//}
		//
		//var fh fs.File
		//var fi os.FileInfo
		//fi, z.Error = z.fs.Stat(fn)
		//if errors.Is(z.Error, os.ErrNotExist) {
		//	// Create new file
		//	fh, z.Error = z.fs.Create(fn)
		//} else {
		//	// Open file
		//	fh, z.Error = z.fs.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0664)
		//}

		fh, err := os.OpenFile(filepath.Join(z.RepoDir, fn), os.O_RDWR|os.O_CREATE, 0664)
		if err != nil {
			z.Error = err
			break
		}
		defer fh.Close()

		fmt.Printf("Saved file '%s'\n", fn)
		_, z.Error = fh.Write(data)
		if z.Error != nil {
			break
		}

		// Run git status before adding the file to the worktree
		//fmt.Println(z.worktree.Status())

		// git add $filePath
		_, z.Error = z.worktree.Add(fn)
		if z.Error != nil {
			break
		}

		//// Run git status after the file has been added adding to the worktree
		//fmt.Println(z.worktree.Status())
		//
		//// git commit -m $message
		//msg := fmt.Sprintf("Updated file '%s'", fn)
		//_, z.Error = z.worktree.Commit(msg, &git.CommitOptions{})
		//if z.Error != nil {
		//	break
		//}

	}

	return z.Error
}

func (z *Git) Status() error {

	for range Only.Once {
		if z.IsNotOk() {
			break
		}

		var status git.Status
		status, z.Error = z.worktree.Status()
		if z.Error != nil {
			break
		}

		if status.String() != "" {
			fmt.Printf("Status of Git\n\trepo: %s\n\tdir: %s\n%s\n",
				z.RepoUrl,
				z.RepoDir,
				status.String(),
			)
		}
	}

	return z.Error
}

func (z *Git) Add(path string) error {

	for range Only.Once {
		if z.IsNotOk() {
			break
		}

		if path == "" {
			path = "."
		}

		fmt.Printf("Adding to Git\n\trepo: %s\n\tdir: %s\n", z.RepoUrl, z.RepoDir)

		_, z.Error = z.worktree.Add(path)
		if z.Error != nil {
			break
		}

		z.Error = z.Status()
		if z.Error != nil {
			break
		}
	}

	//PrintError(z.Error)
	return z.Error
}

func (z *Git) Commit(msg string, args ...interface{}) error {

	for range Only.Once {
		if z.IsNotOk() {
			break
		}

		z.Error = z.Add(".")
		if z.Error != nil {
			break
		}

		cn := &object.Signature {
			Name:  os.Getenv("USERNAME"),
			Email: "",
			When:  time.Now(),
		}

		fmt.Printf("Committing Git\n\trepo: %s\n\tdir: %s\n", z.RepoUrl, z.RepoDir)
		// Similar to git commit -m $message
		var ph plumbing.Hash
		msg := fmt.Sprintf(msg, args...)
		ph, z.Error = z.worktree.Commit(msg, &git.CommitOptions{
			All:       false,
			Author:    cn,
			Committer: cn,
			Parents:   nil,
			SignKey:   nil,
		})
		if z.Error != nil {
			break
		}

		// Similar to git show -s
		var obj *object.Commit
		obj, z.Error = z.repo.CommitObject(ph)
		if z.Error != nil {
			break
		}

		if obj.String() != "" {
			fmt.Printf("Status of Git\n\trepo: %s\n\tdir: %s\n%s\n",
				z.RepoUrl,
				z.RepoDir,
				obj.String(),
			)
		}
	}

	//PrintError(z.Error)
	return z.Error
}

func (z *Git) Pull() error {

	for range Only.Once {
		if z.IsNotOk() {
			break
		}

		pk := z.GetSshAuth()
		if z.Error != nil {
			break
		}

		fmt.Printf("Pulling Git\n\trepo: %s\n\tdir: %s\n", z.RepoUrl, z.RepoDir)
		z.Error = z.worktree.Pull(&git.PullOptions {
			RemoteName:        "",
			ReferenceName:     "",
			SingleBranch:      false,
			Depth:             0,
			Auth:              pk,
			RecurseSubmodules: 0,
			Progress:          os.Stdout,
			Force:             false,
			InsecureSkipTLS:   false,
			CABundle:          nil,
		})
		if z.Error.Error() == "already up-to-date" {
			z.Error = nil
			break
		}
		if z.Error != nil {
			break
		}
	}

	//PrintError(z.Error)
	return z.Error
}

func (z *Git) Push() error {

	for range Only.Once {
		if z.IsNotOk() {
			break
		}

		z.Error = z.Commit("Updated")
		if z.Error != nil {
			break
		}

		pk := z.GetSshAuth()
		if z.Error != nil {
			break
		}

		fmt.Printf("Pushing Git\n\trepo: %s\n\tdir: %s\n", z.RepoUrl, z.RepoDir)
		z.Error = z.repo.Push(&git.PushOptions{
			RemoteName:        "",
			RefSpecs:          nil,
			Auth:              pk,
			Progress:          os.Stdout,
			Prune:             false,
			Force:             false,
			InsecureSkipTLS:   false,
			CABundle:          nil,
			RequireRemoteRefs: nil,
		})
		if z.Error != nil {
			break
		}
	}

	//PrintError(z.Error)
	return z.Error
}

func (z *Git) Diff(path string) error {

	for range Only.Once {
		var c []CommitDiffs

		c, z.Error = z.GetDiffs(path)
		if z.Error != nil {
			break
		}

		if len(c) < 2 {
			fmt.Printf("Not enough revisions to compare.\n")
			break
		}

		f1 := fmt.Sprintf("%s-%s", path, c[0].Hash)
		f1, z.Error = WriteTempFile(f1, c[0].Contents)
		if z.Error != nil {
			break
		}

		f2 := fmt.Sprintf("%s-%s", path, c[1].Hash)
		f2, z.Error = WriteTempFile(f2, c[1].Contents)
		if z.Error != nil {
			break
		}

		if z.DiffCmd == "" {
			z.DiffCmd = "tkdiff"
		}
		z.DiffCmd, z.Error = exec.LookPath(z.DiffCmd)

		cmd := exec.Command(z.DiffCmd, f1, f2)

		var out []byte
		out, z.Error = cmd.Output()
		//if z.Error != nil {
		//	break
		//}

		fmt.Printf("# %s\n", cmd.String())

		fmt.Println(string(out))
		if z.Error != nil {
			break
		}
	}

	//PrintError(z.Error)
	return z.Error
}

type CommitDiffs struct {
	Hash string
	Contents string
}

func (z *Git) GetDiffs(path string) ([]CommitDiffs, error) {
	var ret []CommitDiffs

	for range Only.Once {
		if z.IsNotOk() {
			break
		}

		fmt.Printf("Diff Git\n\trepo: %s\n\tdir: %s\n", z.RepoUrl, z.RepoDir)
		ref, _ := z.repo.Head()
		//fmt.Printf("ref '%s'\n", ref.String())

		commit, _ := z.repo.CommitObject(ref.Hash())
		//fmt.Printf("commit '%s'\n", commit.String())

		var comm []*object.Commit
		commitIter, _ := z.repo.Log(&git.LogOptions{From: commit.Hash})

		_ = commitIter.ForEach(func(c *object.Commit) error {
			comm = append(comm, c)
			return nil
		})

		var lastHash string
		//for k := 0; k < len(comm)-1; k++ {
		for k, _ := range comm {
			fmt.Printf("# Commit number[%d]: %s", k, comm[k].Hash)
			f2, _ := comm[k].File(path)
			if f2 == nil {
				fmt.Println(" - no path")
				continue
			}

			fc, _ := f2.Contents()
			hs := GetHash(fc)
			if hs == lastHash {
				fmt.Println(" - no change")
				continue
			}
			lastHash = hs

			ret = append(ret, CommitDiffs{
				Hash:     comm[k].Hash.String(),
				Contents: fc,
			})
			fmt.Println(" - changed")
		}
	}

	return ret, z.Error
}
