@echo off
START report.exe
timeout 20 > NUL
xcopy /q /f "json\prod\*.json" "U:\Charles\jobJson\prod"
timeout 20 > NUL
xcopy /q /f "json\prod\*.json" "oldJson\prod\*.json"
timeout 20 > NUL
del /q "json\prod\*.json"
PAUSE