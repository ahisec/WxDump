package comm

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// 将文件压缩成 ZIP 文件
func ZipFile(srcPath, zipPath string) error {
	// 创建一个 ZIP 文件
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer zipFile.Close()

	// 创建一个新的 zip.Writer 对象
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 打开待压缩的文件
	file, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// 在 ZIP 中创建一个新文件
	zipEntryWriter, err := zipWriter.Create(filepath.Base(srcPath))
	if err != nil {
		return fmt.Errorf("failed to create zip entry: %w", err)
	}

	// 将原始文件内容复制到 ZIP 文件中
	_, err = io.Copy(zipEntryWriter, file)
	if err != nil {
		return fmt.Errorf("failed to copy file contents to zip: %w", err)
	}

	fmt.Printf("Successfully zipped: %s to %s\n", srcPath, zipPath)
	return nil
}
