// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package drive

import (
	"fmt"
	"io"
	"os"
)

func (f *File) ToDesktopEntry(destAbsPath string) (int, error) {
	desktopEntryPath := sepJoin(".", destAbsPath, DesktopExtension)
	handle, err := os.Create(desktopEntryPath)
	if err != nil {
		return 0, err
	}

	defer func() {
		handle.Close()
		chmodErr := os.Chmod(destAbsPath, 0755)

		if chmodErr != nil {
			fmt.Fprintf(os.Stderr, "%s: [desktopEntry]::chmod %v\n", destAbsPath, chmodErr)
		}

		chTimeErr := os.Chtimes(destAbsPath, f.ModTime, f.ModTime)
		if chTimeErr != nil {
			fmt.Fprintf(os.Stderr, "%s: [desktopEntry]::chtime %v\n", destAbsPath, chTimeErr)
		}
	}()

	return f.SerializeAsDesktopEntry(handle, "url")
}

func (f *File) SerializeAsDesktopEntry(w io.Writer, ext string) (int, error) {
	return io.WriteString(w, fmt.Sprintf("[InternetShortcut]\nURL=%s", f.AlternateLink))
}
