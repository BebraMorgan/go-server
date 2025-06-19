#!/bin/bash

OUTPUT="PROJECT_DOC.md"
echo "# Документация проекта" >$OUTPUT

find . -type f -name '*.go' -exec dirname {} \; | sort -u | while read pkgdir; do
	echo "## Пакет: $pkgdir" >>$OUTPUT
	godocdown "$pkgdir" >>$OUTPUT
	echo -e "\n---\n" >>$OUTPUT
done
