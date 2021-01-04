package _476_Subrectangle_Queries

type SubrectangleQueries struct {
	rectangle_r [][]int
}


func Constructor(rectangle [][]int) SubrectangleQueries {
	rez := SubrectangleQueries{rectangle}
	return rez
}


func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
	for i := row1; i <= row2; i++{
		for j := col1; j <= col2; j++{
			this.rectangle_r[i][j] =  newValue
		}
	}
}


func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.rectangle_r[row][col]
}
