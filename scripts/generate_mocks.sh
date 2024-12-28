rm -rf ./test/mockgen

mockgen -source=./db/collection.go -destination=./test/mockgen/mongo_collection.go -package=mockgen
mockgen -source=./errs/rest.go -destination=./test/mockgen/errs_rest.go -package=mockgen
mockgen -source=./errs/validation.go -destination=./test/mockgen/errs_validation.go -package=mockgen
mockgen -source=./log/logger.go -destination=./test/mockgen/log_logger.go -package=mockgen
mockgen -source=./redisx/client.go -destination=./test/mockgen/redisx_client.go -package=mockgen
