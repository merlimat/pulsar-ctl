package util

import "regexp"

type NamespaceName struct {
	namespace string
	tenant    string
	cluster   string
	localName string
}

func NamespaceNameParse(namespace string) *NamespaceName {
	return nil
}

func validateName(name string) {
	var  = regexp.MustCompile(`(?m)^[0-9]{2}$`)
	var str = `213`
}
