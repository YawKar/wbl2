#!/bin/bash

TESTDATA="testdata"
S_DIR="./sorted_samples"
RS_DIR="./reverse_sorted_samples"
US_DIR="./unique_sorted_samples"
NS_DIR="./numeric_sorted_samples"
RUS_DIR="./reverse_unique_sorted_samples"
RNS_DIR="./reverse_numeric_sorted_samples"
NUS_DIR="./numeric_unique_sorted_samples"
K1S_DIR="./k1_sorted_samples"
K2S_DIR="./k2_sorted_samples"
K3S_DIR="./k3_sorted_samples"
K4S_DIR="./k4_sorted_samples"

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

  rm -rf "$RNS_DIR"
  mkdir "$RNS_DIR"

  rm -rf "$NUS_DIR"
  mkdir "$NUS_DIR"

  rm -rf "$K1S_DIR"
  mkdir "$K1S_DIR"

  rm -rf "$K2S_DIR"
  mkdir "$K2S_DIR"

  rm -rf "$K3S_DIR"
  mkdir "$K3S_DIR"

  rm -rf "$K4S_DIR"
  mkdir "$K4S_DIR"
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

g_RNS() {
  local dirName="$RNS_DIR"
  local sampleFileName="$(basename "$1")"
  local filePath=""$dirName"/"$sampleFileName""
  sort -r -n "$1" > "$filePath"
}

g_NUS() {
  local dirName="$NUS_DIR"
  local sampleFileName="$(basename "$1")"
  local filePath=""$dirName"/"$sampleFileName""
  sort -n -u "$1" > "$filePath"
}

g_K1S() {
  local dirName="$K1S_DIR"
  local sampleFileName="$(basename "$1")"
  local filePath=""$dirName"/"$sampleFileName""
  sort -k 1 "$1" > "$filePath"
}

g_K2S() {
  local dirName="$K2S_DIR"
  local sampleFileName="$(basename "$1")"
  local filePath=""$dirName"/"$sampleFileName""
  sort -k 2 "$1" > "$filePath"
}

g_K3S() {
  local dirName="$K3S_DIR"
  local sampleFileName="$(basename "$1")"
  local filePath=""$dirName"/"$sampleFileName""
  sort -k 3 "$1" > "$filePath"
}

g_K4S() {
  local dirName="$K4S_DIR"
  local sampleFileName="$(basename "$1")"
  local filePath=""$dirName"/"$sampleFileName""
  sort -k 4 "$1" > "$filePath"
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
      g_RNS "$file"
      g_NUS "$file"
      g_K1S "$file"
      g_K2S "$file"
      g_K3S "$file"
      g_K4S "$file"
    fi
done
