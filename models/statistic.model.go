package models

type StatisticQuestCsvRecord struct {
	QuestName     string
	IndicatorName string
	TeamName      string
	FromUser      string
	FromRole      string
	ToUser        string
	ToRole        string
	Value         string
}

// пусть лежит тут, все-таки метод структуры
func (record *StatisticQuestCsvRecord) ToSlice() []string {
	return []string{
		record.QuestName,
		record.IndicatorName,
		record.TeamName,
		record.FromUser,
		record.FromRole,
		record.ToUser,
		record.ToRole,
		record.Value,
	}
}
