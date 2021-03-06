/**
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package realis

type Logger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
	Print(v ...interface{})
}

type NoopLogger struct{}

func (NoopLogger) Printf(format string, a ...interface{}) {}

func (NoopLogger) Print(a ...interface{}) {}

func (NoopLogger) Println(a ...interface{}) {}

type LevelLogger struct {
	Logger
	debug bool
}

func (l *LevelLogger) EnableDebug(enable bool) {
	l.debug = enable
}

func (l LevelLogger) DebugPrintf(format string, a ...interface{}) {
	if l.debug {
		l.Print("[DEBUG] ")
		l.Printf(format, a)
	}
}

func (l LevelLogger) DebugPrint(a ...interface{}) {
	if l.debug {
		l.Print("[DEBUG] ")
		l.Print(a)
	}
}

func (l LevelLogger) DebugPrintln(a ...interface{}) {
	if l.debug {
		l.Print("[DEBUG] ")
		l.Println(a)
	}
}
