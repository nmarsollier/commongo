rm -rf ./test/mockgen

mockgen -source=./pkg/db/collection.go -destination=./test/mockgen/mongo_collection.go -package=mockgen
mockgen -source=./pkg/errs/rest.go -destination=./test/mockgen/errs_rest.go -package=mockgen
mockgen -source=./pkg/errs/validation.go -destination=./test/mockgen/errs_validation.go -package=mockgen
mockgen -source=./pkg/log/logger.go -destination=./test/mockgen/log_logger.go -package=mockgen
mockgen -source=./pkg/rbt/rabbit_channel.go -destination=./test/mockgen/rbt_rabbit.go -package=mockgen
