package tablib_test

import (
	"bytes"
	"encoding/base64"
	"testing"

	tablib "github.com/agrison/go-tablib"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TablibSuite struct{}

var _ = Suite(&TablibSuite{})

func presidentDataset() *tablib.Dataset {
	ds := tablib.NewDataset([]string{"firstName", "lastName", "gpa"})
	ds.AppendValues("John", "Adams", 90)
	ds.AppendValues("George", "Washington", 67)
	ds.AppendValues("Thomas", "Jefferson", 50)
	return ds
}

func presidentDatasetWithTags() *tablib.Dataset {
	ds := tablib.NewDataset([]string{"firstName", "lastName", "gpa"})
	ds.AppendTagged([]interface{}{"John", "Adams", 90}, "Massachusetts")
	ds.AppendTagged([]interface{}{"George", "Washington", 67}, "Virginia")
	ds.AppendTagged([]interface{}{"Thomas", "Jefferson", 50}, "Virginia")
	return ds
}

func frenchPresidentDataset() *tablib.Dataset {
	ds := tablib.NewDataset([]string{"firstName", "lastName", "gpa"})
	ds.AppendValues("Jacques", "Chirac", 88)
	ds.AppendValues("Nicolas", "Sarkozy", 98)
	ds.AppendValues("François", "Hollande", 34)
	return ds
}

func frenchPresidentAdditionalDataset() *tablib.Dataset {
	ds := tablib.NewDataset([]string{"duration", "from"})
	ds.AppendValues(14, "Paris")
	ds.AppendValues(12, "Paris")
	ds.AppendValues(5, "Rouen")
	return ds
}

func carDataset() *tablib.Dataset {
	ds := tablib.NewDataset([]string{"Maker", "Model", "Year"})
	ds.AppendValues("Porsche", "991", 2012)
	ds.AppendValues("Skoda", "Octavia", 2011)
	ds.AppendValues("Ferrari", "458", 2009)
	ds.AppendValues("Citroen", "Picasso II", 2013)
	ds.AppendValues("Bentley", "Continental GT", 2003)
	return ds
}

func validRowAt(d *tablib.Dataset, index int) map[string]interface{} {
	row, _ := d.Row(index)
	return row
}

func lastRow(d *tablib.Dataset) map[string]interface{} {
	row, _ := d.Row(d.Height() - 1)
	return row
}

func (s *TablibSuite) TestDimensions(c *C) {
	ds := presidentDataset()
	c.Assert(ds.Width(), Equals, 3)
	c.Assert(ds.Height(), Equals, 3)
	c.Assert(ds.Headers(), DeepEquals, []string{"firstName", "lastName", "gpa"})
}

func (s *TablibSuite) TestAppendRow(c *C) {
	ds := presidentDataset()
	// too much columns
	c.Assert(ds.AppendValues("a", "b", 50, "d"), Equals, tablib.ErrInvalidDimensions)
	// not enough columns
	c.Assert(ds.AppendValues("a", "b"), Equals, tablib.ErrInvalidDimensions)
	// ok
	c.Assert(ds.AppendValues("foo", "bar", 42), Equals, nil)
	// test values are there
	d := lastRow(ds)
	c.Assert(d["firstName"], Equals, "foo")
	c.Assert(d["lastName"], Equals, "bar")
	c.Assert(d["gpa"], Equals, 42)
}

func (s *TablibSuite) TestAppendColumn(c *C) {
	ds := presidentDataset()
	// too much rows
	c.Assert(ds.AppendColumnValues("foo", "a", "b", "c", "d"), Equals, tablib.ErrInvalidDimensions)
	// not enough columns
	c.Assert(ds.AppendColumnValues("foo", "a", "b"), Equals, tablib.ErrInvalidDimensions)
	// ok
	c.Assert(ds.AppendColumnValues("foo", "a", "b", "c"), Equals, nil)
	// test values are there
	d := ds.Column("foo")
	c.Assert(d[0], Equals, "a")
	c.Assert(d[1], Equals, "b")
	c.Assert(d[2], Equals, "c")
}

func (s *TablibSuite) TestInsert(c *C) {
	ds := presidentDataset()
	// invalid index
	c.Assert(ds.InsertValues(-1, "foo", "bar"), Equals, tablib.ErrInvalidRowIndex)
	c.Assert(ds.InsertValues(100, "foo", "bar"), Equals, tablib.ErrInvalidRowIndex)
	// too much columns
	c.Assert(ds.InsertValues(1, "foo", "bar", 42, "invalid"), Equals, tablib.ErrInvalidDimensions)
	// not enough columns
	c.Assert(ds.InsertValues(1, "foo", "bar"), Equals, tablib.ErrInvalidDimensions)
	// ok
	c.Assert(ds.InsertValues(1, "foo", "bar", 42), Equals, nil)
	// test values are there
	d := validRowAt(ds, 1)
	c.Assert(d["firstName"], Equals, "foo")
	c.Assert(d["lastName"], Equals, "bar")
	c.Assert(d["gpa"], Equals, 42)
}

func (s *TablibSuite) TestInsertColumn(c *C) {
	ds := presidentDataset()
	// invalid index
	c.Assert(ds.InsertColumn(-1, "wut", []interface{}{"foo", "bar"}), Equals, tablib.ErrInvalidColumnIndex)
	c.Assert(ds.InsertColumn(100, "wut", []interface{}{"foo", "bar"}), Equals, tablib.ErrInvalidColumnIndex)
	// too much rows
	c.Assert(ds.InsertColumn(1, "wut", []interface{}{"foo", "bar", "baz", "kidding"}), Equals, tablib.ErrInvalidDimensions)
	// not enough rows
	c.Assert(ds.InsertColumn(1, "wut", []interface{}{"foo", "bar"}), Equals, tablib.ErrInvalidDimensions)
	// ok
	c.Assert(ds.InsertColumn(1, "wut", []interface{}{"foo", "bar", "baz"}), Equals, nil)
	// test values are there
	d := ds.Column("wut")
	c.Assert(d[0], Equals, "foo")
	c.Assert(d[1], Equals, "bar")
	c.Assert(d[2], Equals, "baz")
}

func firstNameB64(row []interface{}) interface{} {
	return base64.StdEncoding.EncodeToString([]byte(row[0].(string)))
}

func lastNameB64(row []interface{}) interface{} {
	return base64.StdEncoding.EncodeToString([]byte(row[1].(string)))
}

func (s *TablibSuite) TestDynamicColumn(c *C) {
	ds := presidentDataset()
	ds.AppendDynamicColumn("firstB64", firstNameB64)
	d := ds.Column("firstB64")
	c.Assert(d[0], Equals, "Sm9obg==") // John
	c.Assert(d[1], Equals, "R2Vvcmdl") // George
	c.Assert(d[2], Equals, "VGhvbWFz") // Thomas

	// invalid index
	c.Assert(ds.InsertDynamicColumn(-1, "foo", lastNameB64), Equals, tablib.ErrInvalidColumnIndex)
	c.Assert(ds.InsertDynamicColumn(100, "foo", lastNameB64), Equals, tablib.ErrInvalidColumnIndex)
	// ok
	c.Assert(ds.InsertDynamicColumn(2, "lastB64", lastNameB64), Equals, nil)
	// check values
	d = ds.Column("lastB64")
	c.Assert(d[0], Equals, "QWRhbXM=")         // Adams
	c.Assert(d[1], Equals, "V2FzaGluZ3Rvbg==") // Washington
	c.Assert(d[2], Equals, "SmVmZmVyc29u")     // Jefferson
}

func (s *TablibSuite) TestRow(c *C) {
	ds := presidentDataset()
	row, err := ds.Row(-1)
	c.Assert(err, Equals, tablib.ErrInvalidRowIndex)
	row, err = ds.Row(100)
	c.Assert(err, Equals, tablib.ErrInvalidRowIndex)
	row, err = ds.Row(1)
	c.Assert(err, Equals, nil)
	c.Assert(row["firstName"], Equals, "George")
	c.Assert(row["lastName"], Equals, "Washington")
}

func (s *TablibSuite) TestRows(c *C) {
	ds := presidentDataset()
	rows, err := ds.Rows(-1, 5)
	c.Assert(err, Equals, tablib.ErrInvalidRowIndex)
	rows, err = ds.Rows(0, 1, 100)
	c.Assert(err, Equals, tablib.ErrInvalidRowIndex)
	rows, err = ds.Rows(1, 2)
	c.Assert(err, Equals, nil)
	c.Assert(rows[0]["firstName"], Equals, "George")
	c.Assert(rows[0]["lastName"], Equals, "Washington")
	c.Assert(rows[1]["firstName"], Equals, "Thomas")
	c.Assert(rows[1]["lastName"], Equals, "Jefferson")
}

func (s *TablibSuite) TestSlice(c *C) {
	ds := presidentDataset()
	_, err := ds.Slice(-1, 5) // invalid lower bound
	c.Assert(err, Equals, tablib.ErrInvalidRowIndex)
	_, err = ds.Slice(0, 100) // invalider upper bound
	c.Assert(err, Equals, tablib.ErrInvalidRowIndex)
	_, err = ds.Slice(1, 0) // lower bound > upper bound
	c.Assert(err, Equals, tablib.ErrInvalidRowIndex)
	s1, err := ds.Slice(1, 2) // single row
	c.Assert(err, Equals, nil)
	c.Assert(s1.Height(), Equals, 1)
	row, _ := s1.Row(0)
	c.Assert(row["firstName"], Equals, "George")
	c.Assert(row["lastName"], Equals, "Washington")
	s2, err := ds.Slice(1, 3) // two rows
	c.Assert(err, Equals, nil)
	c.Assert(s2.Height(), Equals, 2)
	row, _ = s2.Row(0)
	c.Assert(row["firstName"], Equals, "George")
	c.Assert(row["lastName"], Equals, "Washington")
	row, _ = s2.Row(1)
	c.Assert(row["firstName"], Equals, "Thomas")
	c.Assert(row["lastName"], Equals, "Jefferson")
}

func (s *TablibSuite) TestTranspose(c *C) {
	tr := carDataset().Transpose()
	c.Assert(tr.Headers()[0], Equals, "Maker")
	c.Assert(tr.Headers()[1], Equals, "Porsche")
	c.Assert(tr.Headers()[2], Equals, "Skoda")
	c.Assert(len(tr.Headers()), Equals, 6)
	r := validRowAt(tr, 0) // Model
	c.Assert(r["Porsche"], Equals, "991")
	c.Assert(r["Bentley"], Equals, "Continental GT")
	r = validRowAt(tr, 1) // Year
	c.Assert(r["Porsche"], Equals, 2012)
	c.Assert(r["Skoda"], Equals, 2011)
	c.Assert(r["Bentley"], Equals, 2003)
}

func (s *TablibSuite) TestStack(c *C) {
	ds, _ := presidentDataset().Stack(frenchPresidentDataset())
	d := ds.Column("lastName")
	c.Assert(d[0], Equals, "Adams")
	c.Assert(d[1], Equals, "Washington")
	c.Assert(d[2], Equals, "Jefferson")
	c.Assert(d[3], Equals, "Chirac")
	c.Assert(d[4], Equals, "Sarkozy")
	c.Assert(d[5], Equals, "Hollande")

	// check invalid dimensions
	x := frenchPresidentDataset()
	x.DeleteColumn("lastName")
	ds, err := presidentDataset().Stack(x)
	c.Assert(err, Equals, tablib.ErrInvalidDimensions)
}

func (s *TablibSuite) TestStackColumn(c *C) {
	ds, _ := frenchPresidentDataset().StackColumn(frenchPresidentAdditionalDataset())
	d := lastRow(ds)
	c.Assert(d["firstName"], Equals, "François")
	c.Assert(d["lastName"], Equals, "Hollande")
	c.Assert(d["from"], Equals, "Rouen")
	c.Assert(d["duration"], Equals, 5)

	// check invalid dimensions
	x := frenchPresidentAdditionalDataset()
	x.DeleteRow(x.Height() - 1)
	ds, err := frenchPresidentDataset().StackColumn(x)
	c.Assert(err, Equals, tablib.ErrInvalidDimensions)
}

func (s *TablibSuite) TestFiltering(c *C) {
	ds := presidentDatasetWithTags()
	df := ds.Filter("Massachusetts")
	c.Assert(df.Height(), Equals, 1)
	r := lastRow(df)
	c.Assert(r["firstName"], Equals, "John")
	c.Assert(r["lastName"], Equals, "Adams")

	df = ds.Filter("Virginia")
	c.Assert(df.Height(), Equals, 2)
	r = validRowAt(df, 0)
	c.Assert(r["firstName"], Equals, "George")
	c.Assert(r["lastName"], Equals, "Washington")
	r = validRowAt(df, 1)
	c.Assert(r["firstName"], Equals, "Thomas")
	c.Assert(r["lastName"], Equals, "Jefferson")

	df = ds.Filter("Woot")
	c.Assert(df.Height(), Equals, 0)
	c.Assert(df.Width(), Equals, 3)

	df = ds.Filter("Virginia")
	tags, _ := df.Tags(1) // Jefferson
	c.Assert(len(tags), Equals, 1)
	c.Assert(tags[0], Equals, "Virginia")
}

func (s *TablibSuite) TestSort(c *C) {
	ds := presidentDataset().Sort("gpa")
	c.Assert(ds.Height(), Equals, 3)

	r := validRowAt(ds, 0)
	c.Assert(r["firstName"], Equals, "Thomas")
	c.Assert(r["lastName"], Equals, "Jefferson")
	c.Assert(r["gpa"], Equals, 50)

	r = validRowAt(ds, 1)
	c.Assert(r["firstName"], Equals, "George")
	c.Assert(r["lastName"], Equals, "Washington")
	c.Assert(r["gpa"], Equals, 67)

	r = validRowAt(ds, 2)
	c.Assert(r["firstName"], Equals, "John")
	c.Assert(r["lastName"], Equals, "Adams")
	c.Assert(r["gpa"], Equals, 90)

	ds = ds.SortReverse("lastName")
	c.Assert(ds.Height(), Equals, 3)

	r = validRowAt(ds, 0)
	c.Assert(r["firstName"], Equals, "George")
	c.Assert(r["lastName"], Equals, "Washington")

	r = validRowAt(ds, 1)
	c.Assert(r["firstName"], Equals, "Thomas")
	c.Assert(r["lastName"], Equals, "Jefferson")

	r = validRowAt(ds, 2)
	c.Assert(r["firstName"], Equals, "John")
	c.Assert(r["lastName"], Equals, "Adams")
}

func mustBeYoung(val interface{}) bool {
	return val.(int) <= 50
}

func mustBeOld(val interface{}) bool {
	return val.(int) >= 50
}

func (s *TablibSuite) TestValidFailFast(c *C) {
	ds := presidentDataset()

	c.Assert(ds.HasAnyConstraint(), Equals, false)

	ds.ConstrainColumn("gpa", mustBeYoung)
	c.Assert(ds.HasAnyConstraint(), Equals, true)
	c.Assert(ds.ValidFailFast(), Equals, false)

	ds.ConstrainColumn("gpa", mustBeOld)
	c.Assert(ds.ValidFailFast(), Equals, true)
}

func (s *TablibSuite) TestValid(c *C) {
	ds := presidentDataset()

	c.Assert(ds.HasAnyConstraint(), Equals, false)

	ds.ConstrainColumn("gpa", mustBeOld)

	c.Assert(ds.HasAnyConstraint(), Equals, true)
	c.Assert(ds.Valid(), Equals, true)
	c.Assert(len(ds.ValidationErrors), Equals, 0)

	ds.ConstrainColumn("gpa", mustBeYoung)
	c.Assert(ds.Valid(), Equals, false)
	c.Assert(len(ds.ValidationErrors), Equals, 2)

	c.Assert(ds.ValidationErrors[0].Row, Equals, 0)
	c.Assert(ds.ValidationErrors[0].Column, Equals, 2)

	c.Assert(ds.ValidationErrors[1].Row, Equals, 1)
	c.Assert(ds.ValidationErrors[1].Column, Equals, 2)
}

func (s *TablibSuite) TestValidSubset(c *C) {
	ds := presidentDatasetWithTags()

	c.Assert(ds.Valid(), Equals, true)
	c.Assert(ds.ValidSubset(), Equals, ds)

	ds.ConstrainColumn("gpa", mustBeYoung)
	c.Assert(ds.Valid(), Equals, false)
	c.Assert(len(ds.ValidationErrors), Equals, 2)

	// Height is 1
	df := ds.ValidSubset()
	c.Assert(df.Height(), Equals, 1)
	c.Assert(df.HasAnyConstraint(), Equals, false)
	r := validRowAt(df, 0)
	c.Assert(r["firstName"], Equals, "Thomas")
	tags, _ := df.Tags(0)
	c.Assert(tags[0], Equals, "Virginia")
}

func (s *TablibSuite) TestInvalidSubset(c *C) {
	ds := presidentDatasetWithTags()

	c.Assert(ds.Valid(), Equals, true)
	c.Assert(ds.InvalidSubset(), Equals, ds)

	ds.ConstrainColumn("gpa", mustBeYoung)
	c.Assert(ds.Valid(), Equals, false)
	c.Assert(len(ds.ValidationErrors), Equals, 2)

	// Height is 2
	df := ds.InvalidSubset()
	c.Assert(df.Height(), Equals, 2)
	c.Assert(df.HasAnyConstraint(), Equals, false)
	r := validRowAt(df, 0)
	c.Assert(r["firstName"], Equals, "John")
	tags, _ := df.Tags(0)
	c.Assert(tags[0], Equals, "Massachusetts")
	r = validRowAt(df, 1)
	c.Assert(r["firstName"], Equals, "George")
}

func (s *TablibSuite) TestJSON(c *C) {
	ds := frenchPresidentDataset()
	j, _ := ds.JSON()
	c.Assert(j.String(), Equals, "[{\"firstName\":\"Jacques\",\"gpa\":88,\"lastName\":\"Chirac\"},{\"firstName\":\"Nicolas\",\"gpa\":98,\"lastName\":\"Sarkozy\"},{\"firstName\":\"François\",\"gpa\":34,\"lastName\":\"Hollande\"}]")
}

func (s *TablibSuite) TestYAML(c *C) {
	ds := frenchPresidentDataset()
	j, _ := ds.YAML()
	c.Assert(j.String(), Equals, `- firstName: Jacques
  gpa: 88
  lastName: Chirac
- firstName: Nicolas
  gpa: 98
  lastName: Sarkozy
- firstName: François
  gpa: 34
  lastName: Hollande
`)
}

func (s *TablibSuite) TestCSV(c *C) {
	ds := frenchPresidentDataset()
	j, _ := ds.CSV()
	c.Assert(j.String(), Equals, `firstName,lastName,gpa
Jacques,Chirac,88
Nicolas,Sarkozy,98
François,Hollande,34
`)
}

func (s *TablibSuite) TestTSV(c *C) {
	ds := frenchPresidentDataset()
	j, _ := ds.TSV()
	c.Assert(j.String(), Equals, `firstName`+"\t"+`lastName`+"\t"+`gpa
Jacques`+"\t"+`Chirac`+"\t"+`88
Nicolas`+"\t"+`Sarkozy`+"\t"+`98
François`+"\t"+`Hollande`+"\t"+`34
`)
}

func (s *TablibSuite) TestHTML(c *C) {
	ds := frenchPresidentDataset()
	j := ds.HTML()
	c.Assert(j.String(), Equals, `<table class="table table-striped">
	<thead>
		<tr>
			<th>firstName</th>
			<th>lastName</th>
			<th>gpa</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>Jacques</td>
			<td>Chirac</td>
			<td>88</td>
		</tr>
		<tr>
			<td>Nicolas</td>
			<td>Sarkozy</td>
			<td>98</td>
		</tr>
		<tr>
			<td>François</td>
			<td>Hollande</td>
			<td>34</td>
		</tr>
	</tbody>
</table>`)
}

func (s *TablibSuite) TestTabular(c *C) {
	ds := frenchPresidentDataset()
	j := ds.Tabular(tablib.TabularGrid)
	c.Assert(j.String(), Equals, `+--------------+-------------+--------+
|    firstName |    lastName |    gpa |
+==============+=============+========+
|      Jacques |      Chirac |     88 |
+--------------+-------------+--------+
|      Nicolas |     Sarkozy |     98 |
+--------------+-------------+--------+
|     François |    Hollande |     34 |
+--------------+-------------+--------+
`)

	j = ds.Tabular(tablib.TabularSimple)
	c.Assert(j.String(), Equals, `--------------  -------------  --------`+"\n"+
		`    firstName       lastName       gpa `+"\n"+
		`--------------  -------------  --------`+"\n"+
		`      Jacques         Chirac        88 `+"\n"+
		"\n"+
		`      Nicolas        Sarkozy        98 `+"\n"+
		"\n"+
		`     François       Hollande        34 `+"\n"+
		`--------------  -------------  --------`+
		"\n")

	j = ds.Tabular(tablib.TabularCondensed)
	c.Assert(j.String(), Equals, `--------------  -------------  --------`+"\n"+
		`    firstName       lastName       gpa `+"\n"+
		`--------------  -------------  --------`+"\n"+
		`      Jacques         Chirac        88 `+"\n"+
		`      Nicolas        Sarkozy        98 `+"\n"+
		`     François       Hollande        34 `+"\n"+
		`--------------  -------------  --------`+
		"\n")

	j = presidentDataset().Tabular(tablib.TabularMarkdown)
	c.Assert(j.String(), Equals, `|     firstName   |       lastName    |    gpa  |`+" \n"+
		`| --------------  | ---------------   | ------- |`+" \n"+
		`|          John   |          Adams    |     90  |`+" \n"+
		`|        George   |     Washington    |     67  |`+" \n"+
		`|        Thomas   |      Jefferson    |     50  |`+" \n")
}

func (s *TablibSuite) TestMySQL(c *C) {
	ds := frenchPresidentDataset()
	j := ds.MySQL("presidents")
	c.Assert(j.String(), Equals, `CREATE TABLE IF NOT EXISTS presidents
(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	firstName VARCHAR(9),
	lastName VARCHAR(8),
	gpa DOUBLE
);

INSERT INTO presidents VALUES(1, 'Jacques', 'Chirac', 88);
INSERT INTO presidents VALUES(2, 'Nicolas', 'Sarkozy', 98);
INSERT INTO presidents VALUES(3, 'François', 'Hollande', 34);

COMMIT;
`)
}

func (s *TablibSuite) TestPostgres(c *C) {
	ds := frenchPresidentDataset()
	j := ds.Postgres("presidents")
	c.Assert(j.String(), Equals, `CREATE TABLE IF NOT EXISTS presidents
(
	id SERIAL PRIMARY KEY,
	firstName TEXT,
	lastName TEXT,
	gpa NUMERIC
);

INSERT INTO presidents VALUES(1, 'Jacques', 'Chirac', 88);
INSERT INTO presidents VALUES(2, 'Nicolas', 'Sarkozy', 98);
INSERT INTO presidents VALUES(3, 'François', 'Hollande', 34);

COMMIT;
`)
}

func (s *TablibSuite) TestLoadDatabookJSON(c *C) {
	var b bytes.Buffer
	b.WriteString(`[
	  {
	    "title": "Cars",
	    "data": [
	      {"Maker":"Bentley","Model":"Continental GT","Year":2003},
	      {"Maker":"Ferrari","Model":"458","Year":2009},
	      {"Maker":"Skoda","Model":"Octavia","Year":2011},
	      {"Maker":"Porsche","Model":"991","Year":2012},
	      {"Maker":"Citroen","Model":"Picasso II","Year":2013}
	    ]
	  },
	  {
	    "title": "Presidents",
	    "data": [
	      {"Age":90,"First name":"John","Last name":"Adams"},
	      {"Age":83,"First name":"Henry","Last name":"Ford"},
	      {"Age":67,"First name":"George","Last name":"Washington"}
	    ]
	  }
	]`)
	db, _ := tablib.LoadDatabookJSON(b.Bytes())
	c.Assert(db.Size(), Equals, 2)
	c.Assert(db.Sheet("Cars").Title(), Equals, "Cars")
	c.Assert(db.Sheet("Cars").Dataset().Height(), Equals, 5)
	c.Assert(db.Sheet("Presidents").Title(), Equals, "Presidents")
	c.Assert(db.Sheet("Presidents").Dataset().Height(), Equals, 3)
}

func (s *TablibSuite) TestLoadDatabookYAML(c *C) {
	var b bytes.Buffer
	b.WriteString(`- data:
  - Maker: Bentley
    Model: Continental GT
    Year: 2003
  - Maker: Ferrari
    Model: "458"
    Year: 2009
  - Maker: Skoda
    Model: Octavia
    Year: 2011
  - Maker: Porsche
    Model: "991"
    Year: 2012
  - Maker: Citroen
    Model: Picasso II
    Year: 2013
  title: Cars
- data:
  - Age: 90
    First name: John
    Last name: Adams
    Maker: Bentley
    Model: Continental GT
    Year: 2003
  - Age: 83
    First name: Henry
    Last name: Ford
    Maker: Ferrari
    Model: "458"
    Year: 2009
  - Age: 67
    First name: George
    Last name: Washington
    Maker: Skoda
    Model: Octavia
    Year: 2011
  title: Presidents
`)
	db, _ := tablib.LoadDatabookYAML(b.Bytes())
	c.Assert(db.Size(), Equals, 2)
	c.Assert(db.Sheet("Cars").Title(), Equals, "Cars")
	c.Assert(db.Sheet("Cars").Dataset().Height(), Equals, 5)
	c.Assert(db.Sheet("Presidents").Title(), Equals, "Presidents")
	c.Assert(db.Sheet("Presidents").Dataset().Height(), Equals, 3)

	c.Assert(len(db.Sheets()), Equals, 2)
	db.Wipe()
	c.Assert(db.Size(), Equals, 0)
}

func (s *TablibSuite) TestLoadCSV(c *C) {
	var b bytes.Buffer
	b.WriteString(`Maker, Model, Year
Bentley, Continental GT, 2003
Ferrari, 458, 2009
Skoda, Octavia, 2011
Porsche, 991, 2012
Citroen, Picasso II, 2013`)
	ds, _ := tablib.LoadCSV(b.Bytes())
	c.Assert(ds.Height(), Equals, 5)
	c.Assert(ds.Width(), Equals, 3)
	r, _ := ds.Row(0)
	c.Assert(r["Maker"], Equals, "Bentley")
}

func (s *TablibSuite) TestLoadTSV(c *C) {
	var b bytes.Buffer
	b.WriteString(`Maker` + "\t" + `Model` + "\t" + `Year
Bentley` + "\t" + `Continental GT` + "\t" + `2003
Ferrari` + "\t" + `458` + "\t" + `2009
Skoda` + "\t" + `Octavia` + "\t" + `2011
Porsche` + "\t" + `991` + "\t" + `2012
Citroen` + "\t" + `Picasso II` + "\t" + `2013`)
	ds, _ := tablib.LoadTSV(b.Bytes())
	c.Assert(ds.Height(), Equals, 5)
	c.Assert(ds.Width(), Equals, 3)
	r, _ := ds.Row(0)
	c.Assert(r["Maker"], Equals, "Bentley")
}

func (s *TablibSuite) TestLoadXML(c *C) {
	var b bytes.Buffer
	b.WriteString(`<dataset>
 <row>
   <age>90</age>
   <firstName>John</firstName>
   <lastName>Adams</lastName>
 </row>  <row>
   <age>67</age>
   <firstName>George</firstName>
   <lastName>Washington</lastName>
 </row>  <row>
   <age>83</age>
   <firstName>Henry</firstName>
   <lastName>Ford</lastName>
 </row>
</dataset>`)
	ds, _ := tablib.LoadXML(b.Bytes())
	c.Assert(ds.Height(), Equals, 3)
	c.Assert(ds.Width(), Equals, 3)
	r, _ := ds.Row(0)
	c.Assert(r["firstName"], Equals, "John")
	r, _ = ds.Row(1)
	c.Assert(r["lastName"], Equals, "Washington")
}

func (s *TablibSuite) TestXML(c *C) {
	ds := presidentDataset()
	xml, err := ds.XML()
	c.Assert(err, Equals, nil)
	c.Assert(xml.String(), Equals, "<dataset>\n"+
		"  <row>\n"+
		"    <firstName>John</firstName>\n"+
		"    <gpa>90</gpa>\n"+
		"    <lastName>Adams</lastName>\n"+
		"  </row>  <row>\n"+
		"    <firstName>George</firstName>\n"+
		"    <gpa>67</gpa>\n"+
		"    <lastName>Washington</lastName>\n"+
		"  </row>  <row>\n"+
		"    <firstName>Thomas</firstName>\n"+
		"    <gpa>50</gpa>\n"+
		"    <lastName>Jefferson</lastName>\n"+
		"  </row>\n"+
		"  </dataset>")
}

// ---------- Benchmarking ----------

func (s *TablibSuite) BenchmarkAppendRow(c *C) {
	benchDataset1 := frenchPresidentDataset()
	for i := 0; i < c.N; i++ {
		benchDataset1.AppendValues("foo", "bar", 42)
	}
}
