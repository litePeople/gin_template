package dto

// PageAndSize 页码和大小
type PageAndSize struct {
	// 页码
	Page int64 `form:"page" validate:"gte=1" minimum:"1"`
	// 每页大小
	PageSize int64 `form:"pageSize" validate:"gte=1" minimum:"1"`
}
