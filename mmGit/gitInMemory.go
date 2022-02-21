package mmGit

import (
	"GoSungrow/Only"
	"fmt"
	memfs "github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	memory "github.com/go-git/go-git/v5/storage/memory"
	"os"

	//"net/http"
)


func (z *Git) MemConnect() error {

	for range Only.Once {
		if z.IsNotOk() {
			break
		}

		//auth := &http.BasicAuth {
		//	ApiUsername: z.ApiUsername,
		//	ApiPassword: z.ApiPassword,
		//}

		var auth ssh.AuthMethod
		auth, z.Error = ssh.DefaultAuthBuilder("admin-mickh")
		if z.Error != nil {
			break
		}

		//if z.Error = r.Push(&git.PushOptions{Auth: sshAuth}); err != nil {
		//	log.Error().Err(err).Msg("Push err")
		//	os.Exit(1)
		//}
		//if z.Error = r.Push(&git.PushOptions{Auth: sshAuth}); err != nil {
		//	log.Error().Err(err).Msg("Push err")
		//	os.Exit(1)
		//}

		//s := fmt.Sprintf("%s/.ssh/id_rsa", os.Getenv("HOME"))
		//sshKey, err = ioutil.ReadFile(s)
		//signer, err := ssh.ParsePrivateKey([]byte(sshKey))
		//auth = &gitssh.PublicKeys{User: "git", Signer: signer}

		z.storer = memory.NewStorage()
		z.fs = memfs.New()

		z.repo, z.Error = git.Clone(z.storer, z.fs, &git.CloneOptions {
			URL:  z.RepoUrl,
			Auth: auth,
		})
		if z.Error != nil {
			break
		}
	}

	return z.Error
}

func (z *Git) MemSaveFile(fn string, data []byte) error {

	for range Only.Once {
		if z.IsNotOk() {
			break
		}

		z.worktree, z.Error = z.repo.Worktree()
		if z.Error != nil {
			break
		}

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

		fh, err := z.fs.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0664)
		if err != nil {
			z.Error = err
			break
		}
		defer fh.Close()

		_, z.Error = fh.Write(data)
		if z.Error != nil {
			break
		}

		// Run git status before adding the file to the worktree
		fmt.Println(z.worktree.Status())

		// git add $filePath
		_, z.Error = z.worktree.Add(fn)
		if z.Error != nil {
			break
		}

		// Run git status after the file has been added adding to the worktree
		fmt.Println(z.worktree.Status())

		// git commit -m $message
		msg := fmt.Sprintf("Updated file '%s'", fn)
		_, z.Error = z.worktree.Commit(msg, &git.CommitOptions{})
		if z.Error != nil {
			break
		}

	}

	return z.Error
}

func (z *Git) MemCommit(msg string, args ...interface{}) error {

	for range Only.Once {
		if z.IsNotOk() {
			break
		}

		// Similar to git commit -m $message
		msg := fmt.Sprintf(msg, args...)
		_, z.Error = z.worktree.Commit(msg, &git.CommitOptions{})
		if z.Error != nil {
			break
		}

	}

	return z.Error
}
