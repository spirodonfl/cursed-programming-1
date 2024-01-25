#!/bin/bash

# Function to delete a file if it exists
delete_if_exists() {
	if [ -e "$1" ]; then
		rm "$1"
	fi
}

# Function to install TypeScript and npm dependencies
install_dependencies() {
	echo "Installing TypeScript and npm dependencies..."
	npm install -g typescript
	npm install
}

# Delete existing files because we're going to recreate them
delete_if_exists "rust_quine.rs"
delete_if_exists "rust_quine"
delete_if_exists "typescript_quine.ts"
delete_if_exists "python_quine.py"
delete_if_exists "quine.js"

# Install dependencies
install_dependencies

# Loop through the steps
for step in {1..5}; do
	case $step in
	1)
		# Step 1: Compile TypeScript file
		delete_if_exists "quine.js"
		tsc quine.ts
		;;

	2)
		# Step 2: Execute TypeScript output
		node quine.js
		;;

	3)
		# Step 3: Compile Rust file
		delete_if_exists "rust_quine"
		rustc rust_quine.rs
		;;

	4)
		# Step 4: Execute Rust output
		./rust_quine
		;;

	5)
		# Step 5: Execute Python script
		python python_quine.py
		;;
	esac
done
