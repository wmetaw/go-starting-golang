package main

import (
	"fmt"
	"time"
)

func main() {

	// tはtime.Time型
	t := time.Now()
	fmt.Println(t) // 2018-03-03 17:44:13.936106 +0900 JST m=+0.000309812

	t = time.Date(2018, 3, 3, 17, 44, 13, 0, time.Local)
	fmt.Println(t)              // 2018-03-03 17:44:13 +0900 JST
	fmt.Println(t.Year())       // 2018
	fmt.Println(t.YearDay())    // 62 通算日(1~366)
	fmt.Println(t.Month())      // March
	fmt.Println(t.Weekday())    // Saturday
	fmt.Println(t.Day())        // 3
	fmt.Println(t.Hour())       // 17
	fmt.Println(t.Minute())     // 44
	fmt.Println(t.Second())     // 13
	fmt.Println(t.Nanosecond()) // 0
	fmt.Println(t.Zone())       // JST 32400
	fmt.Println("----------------------------------------")

	// 月は1からの連番
	if time.January == 1 {
		// time.goパッケージにて定数定義しているが、String()を実装している為、Printlnの出力は数値ではない
		fmt.Println(time.January)  // January
		fmt.Println(time.December) // December
	}

	// 曜日は日曜日から0スタートの連番
	if time.Sunday == 0 {
		fmt.Println(time.Sunday)
		fmt.Println(time.Saturday)
	}
	fmt.Println("----------------------------------------")

	// 文字列をパースしてtime.Time型を生成
	tt, _ := time.ParseDuration("1h30m")
	fmt.Println(tt.Minutes()) // 90

	// 時間の加算
	t = time.Date(2017, 12, 31, 23, 59, 59, 0, time.Local)
	fmt.Println(t) // 2017-12-31 23:59:59 +0900 JST

	t = t.Add(1 * time.Second)
	fmt.Println(t) // 2018-01-01 00:00:00 +0900 JST

	// 時刻の差分
	t = time.Date(2018, 12, 31, 23, 59, 59, 999999999, time.Local)
	sub := t.Sub(time.Now())
	fmt.Println(sub) // 今年が終わるまであと 7277h30m46.608674999s

	// 「 time.Since 」は「 time.Now().Sub(t) 」のラッパー
	sub = time.Since(t) * -1
	fmt.Println(sub) // 今年が終わるまであと 7277h30m46.608674999s

	fmt.Println("----------------------------------------")

	// 年月日を増減する
	t = time.Date(2018, 12, 31, 23, 59, 59, 999999999, time.Local)
	t = t.AddDate(0, 0, 1) // AddDate(y,m,d)
	fmt.Println(t)         // 2019-01-01 23:59:59.999999999 +0900 JST

	y := time.Now().Year()
	m := time.Now().Month()

	// 今月末を取得
	tmp := time.Date(y, m, 1, 23, 59, 59, 999999999, time.Local).AddDate(0, 1, -1)
	fmt.Println(tmp)

	// 関数化
	fmt.Println(GetLastDate(y, int(m)))

	// 3ヶ月前の月末を取得
	tmp = time.Date(y, m, 1, 23, 59, 59, 999999999, time.Local).AddDate(0, -2, -1)
	fmt.Println(tmp)

	// ex 先月末を取得 dayを0にすると先月末の最終日が返る（危険な香り）
	tmp = time.Date(y, m, 0, 23, 59, 59, 999999999, time.Local)
	fmt.Println(tmp)

	fmt.Println("----------------------------------------")

	// 時刻の比較
	Jan := time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local)
	Feb := time.Date(2018, 2, 1, 0, 0, 0, 0, time.Local)
	fmt.Println(Jan.Before(Feb)) // true  (1月は2月よりも前か?)
	fmt.Println(Jan.After(Feb))  // false (1月は2月よりも後か?)

	fmt.Println("----------------------------------------")

	/* 文字列から時刻を生成
	*
	* ※ UTCになるので注意
	* ※ 2006年01月02日が時刻のフォーマットを表している
	 */
	t, _ = time.Parse("2006/01/02", "2018/03/03")
	fmt.Println(t)
	t, _ = time.Parse("2006年01月02日 15時04分05秒", "2018年03月03日 20時01分30秒")
	fmt.Println(t)
	fmt.Println(t.Local())

	fmt.Println("----------------------------------------")

	// 時刻から文字列を生成
	t = time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.RFC3339Nano))
	fmt.Println(t.Format("2006/01/02"))
	fmt.Println(t.Format("2006年01月02日 15時04分05秒"))

	fmt.Println("----------------------------------------")

	// UNIX時間との相互変換
	t = time.Now()
	fmt.Println(t.Unix())
	fmt.Println(t.UnixNano())

	// UNIXからtime.Timeに変換
	t = time.Unix(t.Unix(), 0)
	fmt.Println(t)
}

// 月末を求める
func GetLastDate(year, month int) time.Time {
	return time.Date(year, time.Month(month+1), 1, 23, 59, 59, 999999999, time.Local).AddDate(0, 0, -1)
}
