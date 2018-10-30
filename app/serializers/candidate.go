package serializers

// Candidate serializer
type Candidate struct {
	FirstName string `json:"firstName" binding:"required,max=50,notspace"`
	LastName  string `json:"lastName" binding:"required,max=50,notspace"`
	Age       uint   `json:"age" binding:"required,min=0"`
	Politics  string `json:"politics" binding:"required,max=500,notspace"`
}
