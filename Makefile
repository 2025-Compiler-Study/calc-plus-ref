SHELL := /bin/bash

PROJECT_DIR=$(shell pwd)
GRAMMAR_DIR=$(PROJECT_DIR)/internal/grammar
GO_PARSER_DIR=$(PROJECT_DIR)/internal/parser
PY_PARSER_DIR=$(PROJECT_DIR)/python

.PHONY: parser

golang_parser:
	@echo "Generate Golang Lexer & Parser code"
	@cd $(GRAMMAR_DIR) && \
		antlr -Dlanguage=Go -o $(GO_PARSER_DIR) -package parser -visitor -listener CalcPlus.g4

python_parser:
	@echo "Generate Python Lexer & Parser code"
	@cd $(GRAMMAR_DIR) && \
		antlr -Dlanguage=Python3 -o $(PY_PARSER_DIR) -visitor -listener CalcPlus.g4