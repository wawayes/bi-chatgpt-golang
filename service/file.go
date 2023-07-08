package service

//func ValidFile(multipartFile *multipart.FileHeader, fileUploadBizEnum requests.FileUploadBizEnum) {
//	// 文件大小
//	fileSize := multipartFile.Size
//	// 文件后缀
//	fileSuffix := strings.TrimPrefix(filepath.Ext(multipartFile.Filename), ".")
//
//	const OneMax = 1024 * 1024
//	if fileUploadBizEnum == requests.Files {
//		if fileSize > OneMax {
//			panic("文件大小不能超过 1M")
//		}
//		allowedFileTypes := []string{"csv", "xlsx"}
//		if !Contains(allowedFileTypes, fileSuffix) {
//			panic("文件类型错误")
//		}
//	}
//}

//func Contains(slice []string, str string) bool {
//	for _, s := range slice {
//		if s == str {
//			return true
//		}
//	}
//	return false
//}
