package crud

import (
	"fmt"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

func ApplyFilter(filter string, tx *gorm.DB) (*gorm.DB, error) {
	// 去除首尾空白和括号
	filter = strings.TrimSpace(filter)
	if len(filter) >= 2 && filter[0] == '(' && filter[len(filter)-1] == ')' {
		filter = filter[1 : len(filter)-1]
	}

	// 分割逻辑运算符
	var conditions []string
	var logicOp string // "AND" 或 "OR"

	// 检查逻辑运算符类型
	if strings.Contains(filter, "&&") {
		conditions = strings.Split(filter, "&&")
		logicOp = "AND"
	} else if strings.Contains(filter, "||") {
		conditions = strings.Split(filter, "||")
		logicOp = "OR"
	} else {
		// 单个条件
		conditions = []string{filter}
		logicOp = ""
	}

	// 遍历所有条件
	for _, cond := range conditions {
		cond = strings.TrimSpace(cond)
		if cond == "" {
			continue
		}

		// 解析字段、操作符、值
		field, op, value, err := parseCondition(cond)
		if err != nil {
			return nil, fmt.Errorf("invalid condition '%s': %v", cond, err)
		}

		// 构建GORM查询条件
		query, arg := buildGormCondition(field, op, value)
		if query == "" {
			return nil, fmt.Errorf("unsupported operator '%s' in condition '%s'", op, cond)
		}

		// 应用条件到查询
		switch logicOp {
		case "OR":
			tx = tx.Or(query, arg)
		default: // 默认为AND
			tx = tx.Where(query, arg)
		}
	}

	return tx, nil
}

// 解析单个条件，返回字段、操作符、值
func parseCondition(cond string) (field, operator, value string, err error) {
	re := regexp.MustCompile(`^(\w+)(~|!~|>=|<=|!=|=|>|<)(.*?)$`)
	matches := re.FindStringSubmatch(cond)
	if len(matches) != 4 {
		return "", "", "", fmt.Errorf("invalid format")
	}
	return matches[1], matches[2], matches[3], nil
}

// 根据字段、操作符和值生成GORM查询条件
func buildGormCondition(field, operator, value string) (string, interface{}) {
	// 去掉value 中的引号
	value = strings.TrimSpace(value)
	if len(value) >= 2 && (value[0] == '"' || value[0] == '\'') && value[0] == value[len(value)-1] {
		value = value[1 : len(value)-1]
	}

	switch operator {
	case "~": // 包含
		return field + " LIKE ?", "%" + value + "%"
	case "!~": // 不包含
		return field + " NOT LIKE ?", "%" + value + "%"
	case ">":
		return field + " > ?", value
	case "<":
		return field + " < ?", value
	case ">=":
		return field + " >= ?", value
	case "<=":
		return field + " <= ?", value
	case "=":
		return field + " = ?", value
	case "!=":
		return field + " != ?", value
	default:
		return "", nil // 不支持的操作符
	}
}
