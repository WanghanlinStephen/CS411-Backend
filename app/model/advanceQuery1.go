package model

func AdvanceQueryOneSQL() error {
	query := "select count(p.PatientId),p.Region from Patients as p natural join Vaccinations as v where v.Fully_vaccinated=1 group by p.Region limit 15"
	if _, err := Db.Exec(query); err != nil {
		return err
	}
	return nil
}
