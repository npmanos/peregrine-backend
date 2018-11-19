package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Pigmice2733/peregrine-backend/internal/store"
	jwt "github.com/dgrijalva/jwt-go"
)

type contextKey string

const (
	keyRolesContext   contextKey = "pigmice_roles"
	keySubjectContext contextKey = "pigmice_subject"
	keyRealmContext   contextKey = "pigmice_realm"
)

// Claims holds the standard jwt claims plus the pigmice roles claim.
type Claims struct {
	Roles   store.Roles `json:"pigmiceRoles"`
	RealmID int64       `json:"pigmiceRealm"`
	jwt.StandardClaims
}

// GetSubject retrieves the subject (user id) from the http context.
func GetSubject(r *http.Request) (int64, error) {
	contextSubject := r.Context().Value(keySubjectContext)
	if contextSubject == nil {
		return 0, fmt.Errorf("no subject set on context")
	}

	sub, ok := contextSubject.(string)
	if !ok {
		return 0, fmt.Errorf("got invalid type for subject")
	}

	id, err := strconv.ParseInt(sub, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to parse subject as int")
	}

	return id, nil
}

// GetRoles retrieves the roles from the http context.
func GetRoles(r *http.Request) store.Roles {
	contextRoles := r.Context().Value(keyRolesContext)
	if contextRoles == nil {
		return store.Roles{}
	}

	roles, ok := contextRoles.(store.Roles)
	if !ok {
		return store.Roles{}
	}

	return roles
}

// GetRealmID retrieves the ID of the user's realm from the http context.
func GetRealmID(r *http.Request) (int64, error) {
	contextRealm := r.Context().Value(keyRealmContext)
	if contextRealm == nil {
		return 0, fmt.Errorf("no realm set on context")
	}

	realmID, ok := contextRealm.(int64)
	if !ok {
		return 0, fmt.Errorf("got invalid type for realm")
	}
	return realmID, nil
}
