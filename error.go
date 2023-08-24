/*
记录程序执行过程中的错误信息，并进行汇总返回
*/
package errlog

var errP *[]string

/**
 * @description:
 * @param {error} err
 * @return
 */
func AddErr(err string) {
	*errP = append(*errP, err)
}

/**
 * @description: 获取处理中的报错
 * @return []error
 */
func GetErrList() []string {
	return *errP
}
