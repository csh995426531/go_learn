// Copyright (c) 2012 - Cloud Instruments Co., Ltd.
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package seelog

import (
	"fmt"
	"io"
	"testing"
)

// fileWriterTestCase is declared in writers_filewriter_test.go

func createRollingSizeFileWriterTestCase(
	files []string,
	fileName string,
	fileSize int64,
	maxRolls int,
	writeCount int,
	resFiles []string,
	nameMode rollingNameMode,
	archiveType rollingArchiveType,
	archiveExploded bool,
	archivePath string) *fileWriterTestCase {

	return &fileWriterTestCase{files, fileName, rollingTypeSize, fileSize, maxRolls, "", writeCount, resFiles, nameMode, archiveType, archiveExploded, archivePath}
}

func createRollingDatefileWriterTestCase(
	files []string,
	fileName string,
	datePattern string,
	writeCount int,
	resFiles []string,
	nameMode rollingNameMode,
	archiveType rollingArchiveType,
	archiveExploded bool,
	archivePath string) *fileWriterTestCase {

	return &fileWriterTestCase{files, fileName, rollingTypeTime, 0, 0, datePattern, writeCount, resFiles, nameMode, archiveType, archiveExploded, archivePath}
}

func TestShouldArchiveWithTar(t *testing.T) {
	compressionType := compressionTypes[rollingArchiveGzip]

	archiveName := compressionType.rollingArchiveTypeName("log", false)

	if archiveName != "log.tar.gz" {
		t.Fatalf("archive name should be log.tar.gz but got %v", archiveName)
	}
}

func TestRollingFileWriter(t *testing.T) {
	t.Logf("Starting rolling file writer tests")
	NewFileWriterTester(rollingfileWriterTests, rollingFileWriterGetter, t).test()
}

//===============================================================

func rollingFileWriterGetter(testCase *fileWriterTestCase) (io.WriteCloser, error) {
	if testCase.rollingType == rollingTypeSize {
		return NewRollingFileWriterSize(testCase.fileName, testCase.archiveType, testCase.archivePath, testCase.fileSize, testCase.maxRolls, testCase.nameMode, testCase.archiveExploded)
	} else if testCase.rollingType == rollingTypeTime {
		return NewRollingFileWriterTime(testCase.fileName, testCase.archiveType, testCase.archivePath, -1, testCase.datePattern, testCase.nameMode, testCase.archiveExploded, false)
	}

	return nil, fmt.Errorf("incorrect rollingType")
}

//===============================================================
var rollingfileWriterTests = []*fileWriterTestCase{
	createRollingSizeFileWriterTestCase([]string{}, "log.testlog", 10, 10, 1, []string{"log.testlog"}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{}, "log.testlog", 10, 10, 2, []string{"log.testlog", "log.testlog.1"}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{"1.log.testlog"}, "log.testlog", 10, 10, 2, []string{"log.testlog", "1.log.testlog", "2.log.testlog"}, rollingNameModePrefix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{"log.testlog.1"}, "log.testlog", 10, 1, 2, []string{"log.testlog", "log.testlog.2"}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{}, "log.testlog", 10, 1, 2, []string{"log.testlog", "log.testlog.1"}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{"log.testlog.9"}, "log.testlog", 10, 1, 2, []string{"log.testlog", "log.testlog.10"}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{"log.testlog.a", "log.testlog.1b"}, "log.testlog", 10, 1, 2, []string{"log.testlog", "log.testlog.1", "log.testlog.a", "log.testlog.1b"}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{}, `dir/log.testlog`, 10, 10, 1, []string{`dir/log.testlog`}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{}, `dir/log.testlog`, 10, 10, 2, []string{`dir/log.testlog`, `dir/1.log.testlog`}, rollingNameModePrefix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{`dir/dir/log.testlog.1`}, `dir/dir/log.testlog`, 10, 10, 2, []string{`dir/dir/log.testlog`, `dir/dir/log.testlog.1`, `dir/dir/log.testlog.2`}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{`dir/dir/dir/log.testlog.1`}, `dir/dir/dir/log.testlog`, 10, 1, 2, []string{`dir/dir/dir/log.testlog`, `dir/dir/dir/log.testlog.2`}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{}, `./log.testlog`, 10, 1, 2, []string{`log.testlog`, `log.testlog.1`}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{`././././log.testlog.9`}, `log.testlog`, 10, 1, 2, []string{`log.testlog`, `log.testlog.10`}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{"dir/dir/log.testlog.a", "dir/dir/log.testlog.1b"}, "dir/dir/log.testlog", 10, 1, 2, []string{"dir/dir/log.testlog", "dir/dir/log.testlog.1", "dir/dir/log.testlog.a", "dir/dir/log.testlog.1b"}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{}, `././dir/log.testlog`, 10, 10, 1, []string{`dir/log.testlog`}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{}, `././dir/log.testlog`, 10, 10, 2, []string{`dir/log.testlog`, `dir/log.testlog.1`}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{`././dir/dir/log.testlog.1`}, `dir/dir/log.testlog`, 10, 10, 2, []string{`dir/dir/log.testlog`, `dir/dir/log.testlog.1`, `dir/dir/log.testlog.2`}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{`././dir/dir/dir/log.testlog.1`}, `dir/dir/dir/log.testlog`, 10, 1, 2, []string{`dir/dir/dir/log.testlog`, `dir/dir/dir/log.testlog.2`}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{}, `././log.testlog`, 10, 1, 2, []string{`log.testlog`, `log.testlog.1`}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{`././././log.testlog.9`}, `log.testlog`, 10, 1, 2, []string{`log.testlog`, `log.testlog.10`}, rollingNameModePostfix, rollingArchiveNone, false, ""),
	createRollingSizeFileWriterTestCase([]string{"././dir/dir/log.testlog.a", "././dir/dir/log.testlog.1b"}, "dir/dir/log.testlog", 10, 1, 2, []string{"dir/dir/log.testlog", "dir/dir/log.testlog.1", "dir/dir/log.testlog.a", "dir/dir/log.testlog.1b"}, rollingNameModePostfix, rollingArchiveNone, true, ""),
	createRollingSizeFileWriterTestCase([]string{"log.testlog", "log.testlog.1"}, "log.testlog", 10, 1, 2, []string{"log.testlog", "log.testlog.2", "dir/log.testlog.1.zip"}, rollingNameModePostfix, rollingArchiveZip, true, "dir"),
	// ====================
}
