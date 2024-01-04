#!/bin/bash

TESTDATA="testdata"
S_DIR="./sorted_samples"
RS_DIR="./reverse_sorted_samples"
US_DIR="./unique_sorted_samples"
NS_DIR="./numeric_sorted_samples"
RUS_DIR="./reverse_unique_sorted_samples"

checkCWDIsTestData() {
  local cwdName="$(basename "$PWD")"
  if [ "$cwdName" != "$TESTDATA" ]; then
    echo "This script can be executed only in the "$TESTDATA"/ folder"
    exit
  fi
}

cleanAndMkGenDirs() {
  rm -rf "$S_DIR"
  mkdir "$S_DIR"

  rm -rf "$RS_DIR"
  mkdir "$RS_DIR"

  rm -rf "$US_DIR"
  mkdir "$US_DIR"

  rm -rf "$NS_DIR"
  mkdir "$NS_DIR"

  rm -rf "$RUS_DIR"
  mkdir "$RUS_DIR"
}

g_S() {
  local dirName="$S_DIR"
  local sampleFileName="$(basename "$1")"
  local filePath=""$dirName"/"$sampleFileName""
  sort "$1" > "$filePath"
}

g_RS() {
  local dirName="$RS_DIR"
  local sampleFileName="$(basename "$1")"
  local filePath=""$dirName"/"$sampleFileName""
  sort -r "$1" > "$filePath"
}

g_US() {
  local dirName="$US_DIR"
  local sampleFileName="$(basename "$1")"
  local filePath=""$dirName"/"$sampleFileName""
  sort -u "$1" > "$filePath"
}

g_NS() {
  local dirName="$NS_DIR"
  local sampleFileName="$(basename "$1")"
  local filePath=""$dirName"/"$sampleFileName""
  sort -n "$1" > "$filePath"
}

g_RUS() {
  local dirName="$RUS_DIR"
  local sampleFileName="$(basename "$1")"
  local filePath=""$dirName"/"$sampleFileName""
  sort -r -u "$1" > "$filePath"
}

checkCWDIsTestData
cleanAndMkGenDirs

for file in ./samples/*; do
    if [ -f "$file" ]; then
      g_S "$file"
      g_RS "$file"
      g_US "$file"
      g_NS "$file"
      g_RUS "$file"
    fi
done
