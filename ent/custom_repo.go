package ent

import "fmt"

func (r *Repo) GetFullName() string {
	return fmt.Sprintf("%s/%s", r.Namespace, r.Name)
}
