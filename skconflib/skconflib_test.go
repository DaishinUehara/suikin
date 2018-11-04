package skconflib_test

import (
	"os"
	"testing"

	"github.com/DaishinUehara/suikin/skconflib"
)

var conf skconflib.SkConf

func TestGetPassword(t *testing.T) {
	var value string
	var key string
	var testno string
	var err error
	var function string

	// TEST-01(何も設定していない場合、空文字が戻ることを確認)
	testno = "[TEST-01]"
	function = "skconflib.GetPassword"
	t.Logf(testno+"[CALL]:"+function+"(%v)\n", key)
	value = conf.GetPassword(key)
	t.Logf(testno+"[RETURN]:value=%v\n", value)
	if value != "" {
		t.Errorf(testno + "[NG]:" + function + "\n")
	} else {
		t.Logf(testno + "[OK]:" + function + "\n")
	}

	// TEST-02(設定した環境変数が読めていることを確認)
	testno = "[TEST-02]"
	err = os.Setenv("SUIKIN_PASSWD_DB", "abc")
	if err != nil {
		t.Errorf(testno+"[INIT][NG]:error=\n", err)
	} else {
		value = os.Getenv("SUIKIN_PASSWD_DB")
		if value != "abc" {
			t.Errorf(testno+"[INIT][NG]:SUIKIN_PASSWD_DB=\n", value)
		} else {
			key = "DB"
			t.Logf(testno+"[CALL]:"+function+"(%v)\n", key)
			value = conf.GetPassword(key)
			t.Logf(testno+"[RETURN]:value=%v\n", value)
			if value != "abc" {
				t.Errorf(testno + "[NG]:" + function + "\n")
			} else {
				t.Logf(testno + "[OK]:" + function + "\n")
			}
		}
	}

	// TEST-03(環境変数を変更してもSigHupで初期化しないと書き換わらないことを確認)
	testno = "[TEST-03]"
	err = os.Setenv("SUIKIN_PASSWD_DB", "def")
	if err != nil {
		t.Errorf(testno+"[INIT][NG]:error=\n", err)
	} else {
		value = os.Getenv("SUIKIN_PASSWD_DB")
		if value != "def" {
			t.Errorf(testno+"[INIT][NG]:SUIKIN_PASSWD_DB=\n", value)
		} else {
			key = "DB"
			t.Logf(testno+"[CALL]:"+function+"(%v)\n", key)
			value = conf.GetPassword(key)
			t.Logf(testno+"[RETURN]:value=%v\n", value)
			if value != "abc" {
				t.Errorf(testno + "[NG]:" + function + "\n")
			} else {
				t.Logf(testno + "[OK]:" + function + "\n")
			}
		}
	}

	// TEST-04(環境変数を変更した後SigHupで初期化すると書き換わることを確認)
	testno = "[TEST-04]"
	err = os.Setenv("SUIKIN_PASSWD_DB", "ghi")
	if err != nil {
		t.Errorf(testno+"[INIT][NG]:error=\n", err)
	} else {
		value = os.Getenv("SUIKIN_PASSWD_DB")
		if value != "ghi" {
			t.Errorf(testno+"[INIT][NG]:SUIKIN_PASSWD_DB=\n", value)
		} else {
			conf.SigHup()
			key = "DB"
			t.Logf(testno+"[CALL]:"+function+"(%v)\n", key)
			value = conf.GetPassword(key)
			t.Logf(testno+"[RETURN]:value=%v\n", value)
			if value != "ghi" {
				t.Errorf(testno + "[NG]:" + function + "\n")
			} else {
				t.Logf(testno + "[OK]:" + function + "\n")
			}
		}
	}

}

func TestGetConfig(t *testing.T) {
	var value string
	var key string
	var testno string
	var err error
	var function string

	// TEST-01(何も設定していない場合、空文字が戻ることを確認)
	testno = "[TEST-01]"
	function = "skconflib.GetConfig"
	t.Logf(testno+"[CALL]:"+function+"(%v)\n", key)
	value = conf.GetConfig(key)
	t.Logf(testno+"[RETURN]:value=%v\n", value)
	if value != "" {
		t.Errorf(testno + "[NG]:" + function + "\n")
	} else {
		t.Logf(testno + "[OK]:" + function + "\n")
	}

	// TEST-02(設定した環境変数が読めていることを確認)
	testno = "[TEST-02]"
	err = os.Setenv("SUIKIN_CONFIG_SERVER_NAME", "server1")
	if err != nil {
		t.Errorf(testno+"[INIT][NG]:error=\n", err)
	} else {
		value = os.Getenv("SUIKIN_CONFIG_SERVER_NAME")
		if value != "server1" {
			t.Errorf(testno+"[INIT][NG]:SUIKIN_CONFIG_SERVER_NAME=\n", value)
		} else {
			key = "SERVER_NAME"
			t.Logf(testno+"[CALL]:"+function+"(%v)\n", key)
			value = conf.GetConfig(key)
			t.Logf(testno+"[RETURN]:value=%v\n", value)
			if value != "server1" {
				t.Errorf(testno + "[NG]:" + function + "\n")
			} else {
				t.Logf(testno + "[OK]:" + function + "\n")
			}
		}
	}

	// TEST-03(環境変数を変更してもSigHupで初期化しないと書き換わらないことを確認)
	testno = "[TEST-03]"
	err = os.Setenv("SUIKIN_CONFIG_SERVER_NAME", "server2")
	if err != nil {
		t.Errorf(testno+"[INIT][NG]:error=\n", err)
	} else {
		value = os.Getenv("SUIKIN_CONFIG_SERVER_NAME")
		if value != "server2" {
			t.Errorf(testno+"[INIT][NG]:SUIKIN_CONFIG_SERVER_NAME=\n", value)
		} else {
			key = "SERVER_NAME"
			t.Logf(testno+"[CALL]:"+function+"(%v)\n", key)
			value = conf.GetConfig(key)
			t.Logf(testno+"[RETURN]:value=%v\n", value)
			if value != "server1" {
				t.Errorf(testno + "[NG]:" + function + "\n")
			} else {
				t.Logf(testno + "[OK]:" + function + "\n")
			}
		}
	}

	// TEST-04(環境変数を変更した後SigHupで初期化すると書き換わることを確認)
	testno = "[TEST-04]"
	err = os.Setenv("SUIKIN_CONFIG_SERVER_NAME", "server3")
	if err != nil {
		t.Errorf(testno+"[INIT][NG]:error=\n", err)
	} else {
		value = os.Getenv("SUIKIN_CONFIG_SERVER_NAME")
		if value != "server3" {
			t.Errorf(testno+"[INIT][NG]:SUIKIN_CONFIG_SERVER_NAME=\n", value)
		} else {
			conf.SigHup()
			key = "SERVER_NAME"
			t.Logf(testno+"[CALL]:"+function+"(%v)\n", key)
			value = conf.GetConfig(key)
			t.Logf(testno+"[RETURN]:value=%v\n", value)
			if value != "server3" {
				t.Errorf(testno + "[NG]:" + function + "\n")
			} else {
				t.Logf(testno + "[OK]:" + function + "\n")
			}
		}
	}
}

func TestGetLogConfig(t *testing.T) {
	var value string
	var key string
	var testno string
	var err error
	var function string

	// TEST-01(何も設定していない場合、空文字が戻ることを確認)
	testno = "[TEST-01]"
	function = "skconflib.GetLogConfig"
	t.Logf(testno+"[CALL]:"+function+"(%v)\n", key)
	value = conf.GetLogConfig(key)
	t.Logf(testno+"[RETURN]:value=%v\n", value)
	if value != "" {
		t.Errorf(testno + "[NG]:" + function + "\n")
	} else {
		t.Logf(testno + "[OK]:" + function + "\n")
	}

	// TEST-02(設定した環境変数が読めていることを確認)
	testno = "[TEST-02]"
	err = os.Setenv("SUIKIN_LOG_FILE", "log1")
	if err != nil {
		t.Errorf(testno+"[INIT][NG]:error=\n", err)
	} else {
		value = os.Getenv("SUIKIN_LOG_FILE")
		if value != "log1" {
			t.Errorf(testno+"[INIT][NG]:SUIKIN_LOG_FILE=\n", value)
		} else {
			key = "FILE"
			t.Logf(testno+"[CALL]:"+function+"(%v)\n", key)
			value = conf.GetLogConfig(key)
			t.Logf(testno+"[RETURN]:value=%v\n", value)
			if value != "log1" {
				t.Errorf(testno + "[NG]:" + function + "\n")
			} else {
				t.Logf(testno + "[OK]:" + function + "\n")
			}
		}
	}

	// TEST-03(環境変数を変更してもSigHupで初期化しないと書き換わらないことを確認)
	testno = "[TEST-03]"
	err = os.Setenv("SUIKIN_LOG_FILE", "log2")
	if err != nil {
		t.Errorf(testno+"[INIT][NG]:error=\n", err)
	} else {
		value = os.Getenv("SUIKIN_LOG_FILE")
		if value != "log2" {
			t.Errorf(testno+"[INIT][NG]:SUIKIN_LOG_FILE=\n", value)
		} else {
			key = "FILE"
			t.Logf(testno+"[CALL]:"+function+"(%v)\n", key)
			value = conf.GetLogConfig(key)
			t.Logf(testno+"[RETURN]:value=%v\n", value)
			if value != "log1" {
				t.Errorf(testno + "[NG]:" + function + "\n")
			} else {
				t.Logf(testno + "[OK]:" + function + "\n")
			}
		}
	}

	// TEST-04(環境変数を変更した後SigHupで初期化すると書き換わることを確認)
	testno = "[TEST-04]"
	err = os.Setenv("SUIKIN_LOG_FILE", "log3")
	if err != nil {
		t.Errorf(testno+"[INIT][NG]:error=\n", err)
	} else {
		value = os.Getenv("SUIKIN_LOG_FILE")
		if value != "log3" {
			t.Errorf(testno+"[INIT][NG]:SUIKIN_LOG_FILE=\n", value)
		} else {
			conf.SigHup()
			key = "FILE"
			t.Logf(testno+"[CALL]:"+function+"(%v)\n", key)
			value = conf.GetLogConfig(key)
			t.Logf(testno+"[RETURN]:value=%v\n", value)
			if value != "log3" {
				t.Errorf(testno + "[NG]:" + function + "\n")
			} else {
				t.Logf(testno + "[OK]:" + function + "\n")
			}
		}
	}

}
