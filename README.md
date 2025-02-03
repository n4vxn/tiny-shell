# Tiny-Shell

## What is a Shell?

A **shell** is a command-line interface that allows users to interact with an operating system by typing commands. It acts as an intermediary between the user and the operating system, interpreting and executing the commands. There are different types of shells, such as **Bash**, **Zsh**, **PowerShell**, and more. These shells provide features like navigation, file manipulation, and running applications, all from a terminal window.

## Project Overview

This project is a simple **custom shell** implemented in **Go** that mimics basic shell functionality. It supports a set of commands to help with navigation, file creation, deletion, and other basic file system tasks.

## Features

- **Change Directory**: `crib <directory>` - Change the current working directory.
- **List Files**: `peek` - List files and directories in the current directory.
- **Print Current Directory**: `whereami` - Print the current directory path.
- **Create Directory**: `cribup <directory>` - Create a new directory.
- **Yell**: `yell <message>` - Print a message to the terminal.
- **Create File**: `new <filename>` - Create a new file.
- **Remove Directory**: `cribdown <directory>` - Remove a directory and its contents.
- **Remove File**: `trash <filename>` - Remove a file.