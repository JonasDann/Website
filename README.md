# Website
Website running on dann.io

The data for the CV is structured according to [CVGen](https://github.com/pschichtel/CVGen) by [Phillip Schichtel](https://github.com/pschichtel). The supported structure is as follows (unless otherwise stated the fields are of type string):
- /cv/activity (extraordinary experience like privat projects)
  - @time
  - @description
- /cv/bio (general data about the person)
  - @first_name
  - @last_name
  - @birth_date
  - @birth_location
  - @street
  - @house_number
  - @postal_code
  - @city
  - @phone
  - @email
  - @links (optional) (JSON object which maps names to a list of links)
- /cv/education (schools and universities)
  - @from
  - @until (optional)
  - @name
  - @location
  - @degree (optional)
  - @grade (optional)
  - @description (optional)
- /cv/experience (professinal experience like companies the person worked for)
  - @from
  - @until (optional)
  - @name
  - @location
  - @position
  - @description (optional)
  - @applied_tech (optional) (list of strings)