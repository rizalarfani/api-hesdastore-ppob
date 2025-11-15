GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

# Live Reload
watch-prepare:
	curl sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh

# run dev
watch:
	bin/air