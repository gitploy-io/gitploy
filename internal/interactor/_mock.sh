mockgen \
    -aux_files \
github.com/gitploy-io/gitploy/internal/interactor=user.go\
,github.com/gitploy-io/gitploy/internal/interactor=repo.go\
,github.com/gitploy-io/gitploy/internal/interactor=deployment.go\
,github.com/gitploy-io/gitploy/internal/interactor=deploymentstatistics.go\
,github.com/gitploy-io/gitploy/internal/interactor=lock.go\
,github.com/gitploy-io/gitploy/internal/interactor=event.go\
    -source ./interface.go \
    -package mock \
    -destination ./mock/pkg.go