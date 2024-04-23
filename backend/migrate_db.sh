#!/bin/bash
migrate -source file://db/migrations -database postgres://postgres:postgres@localhost:5432/investment_warlock?sslmode=disable up