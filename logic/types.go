package logic

// Entry is a struct that holds the data for the html template
type Entry struct {
	Columns []string
	Rows    [][]string
	View    string
}

// Contact is a struct that holds the data for the contact page
type Contact struct {
	ID        int
	Nom       string
	Prenom    string
	Email     string
	Telephone string
}

// Formation is a struct that holds the data for the formations page
type Formation struct {
	ID            int
	Etablissement string
	Diplome       string
	DateDebut     string
	DateFin       string
	Description   string
}

// Experience is a struct that holds the data for the experiences page
type Experience struct {
	ID          int
	Poste       string
	Societe     string
	DateDebut   string
	DateFin     string
	Description string
}

// Competence is a struct that holds the data for the competences page
type Competence struct {
	ID          int
	Nom         string
	Niveau      int
	Description string
}
