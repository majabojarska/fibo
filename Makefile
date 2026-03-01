.PHONY: docs
docs:
	swag fmt
	swag init --dir internal/routes/ --parseInternal --generalInfo router.go
