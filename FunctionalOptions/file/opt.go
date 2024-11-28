package file

import "os"

type Options struct {
	UID         int
	GID         int
	Flags       int
	Contents    string
	Permissions os.FileMode
}

type Option func(*Options)

func WithUID(userID int) Option {
	return func(args *Options) {
		args.UID = userID
	}
}

func WithGID(groupID int) Option {
	return func(args *Options) {
		args.GID = groupID
	}
}

func WithContents(content string) Option {
	return func(args *Options) {
		args.Contents = content
	}
}
func WithFlags(flag int) Option {
	return func(args *Options) {
		args.Flags = flag
	}
}

func WithPermission(perm os.FileMode) Option {
	return func(args *Options) {
		args.Permissions = perm
	}
}
