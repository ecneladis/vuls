/* Vuls - Vulnerability Scanner
Copyright (C) 2016  Future Architect, Inc. Japan.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package report

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path/filepath"

	c "github.com/future-architect/vuls/config"
	"github.com/future-architect/vuls/models"
	"github.com/k0kubun/pp"
)

// LocalFileWriter writes results to file.
type LocalFileWriter struct {
	CurrentDir string
}

func (w LocalFileWriter) Write(rs ...models.ScanResult) (err error) {
	for _, r := range rs {
		path := filepath.Join(w.CurrentDir, r.ReportFileName())

		if c.Conf.FormatJSON {
			p := path + ".json"
			var b []byte
			if b, err = json.Marshal(r); err != nil {
				return fmt.Errorf("Failed to Marshal to JSON: %s", err)
			}
			if err := ioutil.WriteFile(p, b, 0600); err != nil {
				return fmt.Errorf("Failed to write JSON. path: %s, err: %s", p, err)
			}
		}

		if c.Conf.FormatDetailText {
			p := path + ".txt"
			text, err := toPlainText(r)
			if err != nil {
				return err
			}
			if err := ioutil.WriteFile(
				p, []byte(text), 0600); err != nil {
				return fmt.Errorf(
					"Failed to write text files. path: %s, err: %s",
					p, err)
			}
		}

		if c.Conf.FormatSummaryText {
			p := path + ".txt"
			text, err := toPlainText(r)
			if err != nil {
				return err
			}
			if err := ioutil.WriteFile(
				p, []byte(text), 0600); err != nil {
				return fmt.Errorf(
					"Failed to write text files. path: %s, err: %s",
					p, err)
			}
		}

		if c.Conf.FormatXML {
			p := path + ".xml"
			pp.Println(p)
			var b []byte
			if b, err = xml.Marshal(r); err != nil {
				return fmt.Errorf("Failed to Marshal to XML: %s", err)
			}
			allBytes := bytes.Join([][]byte{[]byte(xml.Header + vulsOpenTag), b, []byte(vulsCloseTag)}, []byte{})
			if err := ioutil.WriteFile(p, allBytes, 0600); err != nil {
				return fmt.Errorf("Failed to write XML. path: %s, err: %s", p, err)
			}
		}
	}
	return nil
}
