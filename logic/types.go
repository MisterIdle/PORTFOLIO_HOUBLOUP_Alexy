package logic

type Entry struct {
	Columns []string
	Values  []interface{}
}

type Contact struct {
	ID        int
	Nom       string
	Prenom    string
	Email     string
	Telephone string
}

type Formation struct {
	ID            int
	Etablissement string
	Diplome       string
	DateDebut     string
	DateFin       string
	Description   string
}

type Experience struct {
	ID          int
	Poste       string
	Societe     string
	DateDebut   string
	DateFin     string
	Description string
}

type Competence struct {
	ID          int
	Nom         string
	Niveau      int
	Description string
}
