#!/bin/bash

GO_PATH="$HOME/go/bin"

if [ ! -d "$GO_PATH" ]; then
    echo "Предупреждение: $GO_PATH не существует"
    echo "Но добавим в PATH на будущее"
fi

PATH_LINE="export PATH=\"\$PATH:$GO_PATH\""

add_to_config() {
    local file="$1"
    local line="$2"
    
    [ -f "$file" ] || return
    
    if ! grep -q "$line" "$file"; then
        echo "$line" >> "$file"
    fi
}

add_to_config ~/.bashrc "$PATH_LINE"
add_to_config ~/.zshrc "$PATH_LINE"

go build
go install
