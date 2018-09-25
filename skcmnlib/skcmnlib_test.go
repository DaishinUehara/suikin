package skcmnlib_test

import (
	"testing"

	"github.com/DaishinUehara/suikin/skcmnlib"
)

func TestCammaDivide(t *testing.T) {
	selectColumnName := make([]string, 0, 5)
	selectColumnName = append(selectColumnName, "工場")
	selectColumnName = append(selectColumnName, "売上")
	selectColumnName = append(selectColumnName, "原価")
	selectColumnName = append(selectColumnName, "売上総利益")

	incolumnname, outcolumnname, err := skcmnlib.CammaDivide(selectColumnName)

	if err != nil {
		t.Errorf("skcmnlib.CammaDivide:想定できないエラー: %v", err)
	}

	for i, ans := range selectColumnName {
		if incolumnname[i] != ans {
			t.Errorf("skcmnlib.CammaDivide:incolumnname[%d]=[戻値:%s]:[想定:%s]\n", i, incolumnname[i], ans)
			t.Errorf("skcmnlib.CammaDivide:outcolumnname[%d]=[戻値:%s]:[想定:%s]\n", i, outcolumnname[i], ans)
		}
	}
}
