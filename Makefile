SHELL := /bin/bash

PROJECT_DIR=$(shell pwd)
GRAMMAR_DIR=$(PROJECT_DIR)/internal/grammar
PARSER_DIR=$(PROJECT_DIR)/internal/parser

.PHONY: parser

parser:
	@echo "Generate Lexer & Parser code"
	@cd $(GRAMMAR_DIR) && \
		antlr -Dlanguage=Go -o $(PARSER_DIR) -package parser -visitor -listener CalcPlus.g4