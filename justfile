ANTLR_JAR_NAME := "antlr-4.9-complete.jar"
ANTLR_JAR := "/usr/local/lib/antlr-4.9-complete.jar"
ANTLR_JAR_DOWNLOAD_URL := "https://www.antlr.org/download/antlr-4.9-complete.jar"

default:
	@just --list
antlr:
    cd parser; java -Xmx500M -cp "{{ANTLR_JAR}}:$CLASSPATH" org.antlr.v4.Tool -no-listener -visitor -Dlanguage=Go *.g4
