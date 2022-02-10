package mmGit

import (
	"GoSungro/Only"
	"errors"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)


type Git struct {
	Username string
	Password string
	KeyFile  string
	Token    string
	RepoUrl  string
	RepoDir  string
	DiffCmd  string

	repo   *git.Repository

	modeMemory bool
	worktree *git.Worktree
	storer *memory.Storage
	fs    billy.Filesystem

	Error error
}

func New() *Git {
	var ret Git

	for range Only.Once {
	}

	return &ret
}

func (z *Git) SetAuth(username string, password string) error {

	for range Only.Once {
		if username == "" {
			z.Error = errors.New("username empty")
			break
		}
		z.Username = username

		if password == "" {
			z.Error = errors.New("password empty")
			break
		}
		z.Password = password
	}

	return z.Error
}

func (z *Git) SetKeyFile(path string) error {

	for range Only.Once {
		if path == "" {
			break
		}

		z.Error = checkKeyFile(path)
		if z.Error != nil {
			break
		}

		z.KeyFile = path
	}

	return z.Error
}

func (z *Git) SetToken(t string) error {

	for range Only.Once {
		if t == "" {
			break
		}

		z.Token = t
	}

	return z.Error
}

func (z *Git) SetRepo(repo string) error {

	for range Only.Once {
		if repo == "" {
			z.Error = errors.New("repo empty")
			break
		}
		z.RepoUrl = repo
	}

	return z.Error
}

func (z *Git) SetDir(dir string) error {

	for range Only.Once {
		if dir == "" {
			z.Error = errors.New("dir empty")
			break
		}
		z.RepoDir = dir
	}

	return z.Error
}

func (z *Git) SetDiffCmd(cmd string) error {

	for range Only.Once {
		if cmd == "" {
			cmd = "tkdiff"
		}
		z.DiffCmd = cmd
	}

	return z.Error
}

func (z *Git) IsOk() bool {
	var ok bool

	for range Only.Once {
		//if z.ApiUsername == "" {
		//	z.Error = errors.New("username empty")
		//	break
		//}
		//
		//if z.ApiPassword == "" {
		//	z.Error = errors.New("password empty")
		//	break
		//}

		if z.RepoUrl == "" {
			z.Error = errors.New("repo empty")
			break
		}

		if z.RepoDir == "" {
			z.Error = errors.New("repo dir empty")
			break
		}

		ok = true
	}

	return ok
}
func (z *Git) IsNotOk() bool {
	return !z.IsOk()
}
