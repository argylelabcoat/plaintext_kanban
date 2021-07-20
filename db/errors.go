package db

import "fmt"

type ErrorTaskMissingMeta struct {
	FilePath string
	MetaTag  string
}

func (etmm *ErrorTaskMissingMeta) Error() string {
	return fmt.Sprintf("Task located at: %v missing required meta tag: %v.", etmm.FilePath, etmm.MetaTag)
}

func NewErrorTaskMissingMeta(path, meta string) (err error) {
	return &ErrorTaskMissingMeta{FilePath: path, MetaTag: meta}
}
