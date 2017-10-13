package utils

func NewString(value string) *string {
    result := string(value)
    return &result
}
