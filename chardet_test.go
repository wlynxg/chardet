package chardet

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDetect(t *testing.T) {
	// 遍历 testdata 目录
	err := filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过目录
		if info.IsDir() {
			return nil
		}

		// 获取期望的编码（从父目录名称获取）
		parent := filepath.Base(filepath.Dir(path))
		expectedEncoding := strings.ToUpper(parent) // 转换为大写以统一比较

		// 读取文件内容
		content, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("Failed to read file %s: %v", path, err)
			return nil
		}

		// 调用被测试的函数
		result := Detect(content)

		// 比较结果
		if strings.ToUpper(result.Encoding) != expectedEncoding {
			t.Errorf("File: %s\nExpected encoding: %s\nGot: %s",
				path,
				expectedEncoding,
				result.Encoding,
			)
		}

		return nil
	})

	if err != nil {
		t.Fatalf("Failed to walk testdata directory: %v", err)
	}
}

// TestDetectSubset 测试特定子集
func TestDetectSubset(t *testing.T) {
	testCases := []struct {
		dir      string
		minFiles int // 最小期望文件数
	}{
		// {"Ascii", 1},
		// {"Big5", 1},
		// {"CP932", 1},
		// {"CP949", 1},
		// {"EUC-JP", 1},
		// {"EUC-KR", 1},
		// {"EUC-TW", 1},
		// {"GB2312", 1},
		{"IBM855", 1},
	}

	for _, tc := range testCases {
		t.Run(tc.dir, func(t *testing.T) {
			dirPath := filepath.Join("testdata", tc.dir)
			files, err := os.ReadDir(dirPath)
			if err != nil {
				t.Skipf("Directory %s not found: %v", dirPath, err)
				return
			}

			if len(files) < tc.minFiles {
				t.Errorf("Expected at least %d files in %s, found %d",
					tc.minFiles, dirPath, len(files))
				return
			}

			for _, file := range files {
				if file.IsDir() {
					continue
				}

				filePath := filepath.Join(dirPath, file.Name())
				t.Logf("%s", filePath)
				content, err := os.ReadFile(filePath)
				if err != nil {
					t.Errorf("Failed to read file %s: %v", filePath, err)
					continue
				}

				result := Detect(content)
				if strings.ToUpper(result.Encoding) != strings.ToUpper(tc.dir) {
					t.Errorf("\nFile: %s\nExpected: %s\nGot: %s\nConfidence: %f",
						filePath,
						strings.ToUpper(tc.dir),
						result.Encoding,
						result.Confidence,
					)
				}
			}
		})
	}
}

// TestDetectWithBenchmark 包含性能测试
func TestDetectWithBenchmark(t *testing.T) {
	var totalSize int64
	var fileCount int

	// 遍历所有测试文件并记录性能
	err := filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("Failed to read file %s: %v", path, err)
			return nil
		}

		fileCount++
		totalSize += info.Size()

		// 记录检测时间
		result := Detect(content)

		// 验证结果
		expected := strings.ToUpper(filepath.Base(filepath.Dir(path)))
		if strings.ToUpper(result.Encoding) != expected {
			t.Errorf("File %s: expected %s, got %s", path, expected, result.Encoding)
		}

		return nil
	})

	if err != nil {
		t.Fatalf("Failed to walk testdata directory: %v", err)
	}

	t.Logf("Processed %d files, total size: %d bytes", fileCount, totalSize)
}
