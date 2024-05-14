/*
 * Параллельная обработка gzipper
 *
 * Напишите пакет (gzipper) для сжатия файлов с помощью compress/gzip. Для
 * этого реализуйте следующие функции:
 * - FileNameGen(dir string, pattern *regexp.Regexp) <-chan Work для получения
 *   файлов в директории по заданному регулярному выражению
 * - compress(jobs <-chan Work) для сжатия файлов (на каждый файл отдельная
 *   горутина) из канала и записи результатов на диск.
 * Имена сжатых файлов должны быть сформированы по правилу
 * имяисходногофайла.gz, например для файла myfile.txt -> myfile.txt.gz.
 * Используйте горутины для параллельной обработки файлов.
 *
 * Примечания
 * Код программы должен содержать объявление следующей структуры:
 * type Work struct {
 *   FilePath string
 * }
 *
 */

package gzipper

import (
	"compress/gzip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

type Work struct {
	Filepath string
}

func FileNameGen(dir string, pattern *regexp.Regexp) <-chan Work {
	jobs := make(chan Work)
	go func() {
		defer close(jobs)
		filepath.Walk(
			dir,
			func(path string, finfo fs.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !finfo.IsDir() {
					result := pattern.Find(
						[]byte(filepath.Base(path)),
					)
					if len(result) > 0 {
						jobs <- Work{path}
					}
				}
				return nil
			},
		)
	}()
	return jobs
}

func compress(jobs <-chan Work) {
	for work := range jobs {
		out, err := os.Create(work.Filepath + ".gz")
		if err != nil {
			continue
		}
		defer out.Close()

		gw := gzip.NewWriter(out)
		defer gw.Close()

		in, err := os.Open(work.Filepath)
		if err != nil {
			continue
		}
		defer in.Close()

		_, err = io.Copy(gw, in)
		if err != nil {
			continue
		}
	}
}
