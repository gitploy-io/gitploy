mockgen \
    -aux_files \
github.com/gitploy-io/gitploy/internal/interactor=user.go\
,github.com/gitploy-io/gitploy/internal/interactor=repo.go\
,github.com/gitploy-io/gitploy/internal/interactor=perm.go\
,github.com/gitploy-io/gitploy/internal/interactor=config.go\
,github.com/gitploy-io/gitploy/internal/interactor=deployment.go\
,github.com/gitploy-io/gitploy/internal/interactor=deploymentstatistics.go\
,github.com/gitploy-io/gitploy/internal/interactor=lock.go\
,github.com/gitploy-io/gitploy/internal/interactor=event.go\
,github.com/gitploy-io/gitploy/internal/interactor=review.go\
    -source ./interface.go \
    -package mock \
    -destination ./mock/pkg.go