package moon

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Shell struct {
	e map[string]string
}

type ErrorCode uint8

func (sh *Shell) Env(o io.Writer, e io.Writer) ErrorCode {
	const (
		Success ErrorCode = iota
		ErrPrintEnv
	)
	for k, v := range sh.e {
		fmt.Fprintf(o, "%s=%s\n", k, v)
	}
	return Success
}

func (sh *Shell) Cd(args []string, e io.Writer) ErrorCode {
	const (
		Success ErrorCode = iota
		ErrPWDIsUnset
		ErrCantCreateNewPwd
		ErrDirDoesntExit
		ErrOsInternal
		ErrNotADirectory
		ErrHomeDirIsNotSet
	)
	if len(args) == 0 {
		if homeDir, set := sh.e[E_HOME]; !set {
			fmt.Fprintf(e, "cd: $%s is not set\n", E_HOME)
			return ErrHomeDirIsNotSet
		} else {
			sh.e[E_PWD] = homeDir
		}
		return Success
	}
	dir := args[0]
	pwd, set := sh.e[E_PWD]
	if !set {
		fmt.Fprintf(e, "cd: $%s is unset!\n", E_PWD)
		return ErrPWDIsUnset
	}
	if newPath, err := filepath.Rel(pwd, dir); err != nil {
		fmt.Fprintf(e, "cd: can't change dir to %s from %s\n", pwd, dir)
		return ErrCantCreateNewPwd
	} else if f, err := os.Stat(newPath); err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(e, "cd: %s: no such file or directory\n", newPath)
			return ErrDirDoesntExit
		} else {
			fmt.Fprintf(e, "cd: %s: err: %s\n", newPath, err)
			return ErrOsInternal
		}
	} else if !f.IsDir() {
		fmt.Fprintf(e, "cd: %s: not a directory", newPath)
		return ErrNotADirectory
	} else {
		sh.e[E_PWD] = newPath
		return Success
	}
}

func (sh *Shell) Echo(args []string, o io.Writer, e io.Writer) ErrorCode {
	const (
		Success ErrorCode = iota
	)
	for i, arg := range args {
		if i == 0 {
			fmt.Fprint(o, arg)
		} else {
			fmt.Fprintf(o, " %s", arg)
		}
	}
	fmt.Fprintln(o) // just '\n'
	return Success
}
