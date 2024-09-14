# Makefile

-include .env

RULES_FILE=scripts/make/rules.mk

# APP_CMD_NAMES = \
# 	studio \
# 	end2endtests

# APP_PROTO_FILES= \
# 	api/studio/studio.proto \
# 	api/review/review.proto \


include $(RULES_FILE)


