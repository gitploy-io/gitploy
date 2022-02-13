mockgen \
    -aux_files \
github.com/gitploy-io/gitploy/internal/interactor=user.go\
,github.com/gitploy-io/gitploy/internal/interactor=sync.go\
    -source ./interface.go \
    -package mock \
    -destination ./mock/pkg.go