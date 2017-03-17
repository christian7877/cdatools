package fixtures

var Cms9v4a = []byte(`
{
  "id": "40280381-4DE7-DB4D-014D-E8C552E9025F",
  "nqf_id": "0480",
  "hqmf_id": "40280381-4DE7-DB4D-014D-E8C552E9025F",
  "hqmf_set_id": "7D374C6A-3821-4333-A1BC-4531005D77B8",
  "hqmf_version_number": 4,
  "cms_id": "CMS9v4",
  "name": "Exclusive Breast Milk Feeding",
  "description": "PC-05 Exclusive breast milk feeding during the newborn's entire hospitalization.\n\nPC-05a Exclusive breast milk feeding during the newborn's entire hospitalization considering mother's choice.",
  "type": "eh",
  "category": "Newborn",
  "map_fn": "function() {\n          var patient = this;\n          var effective_date = <%= effective_date %>;\n          var enable_logging = <%= enable_logging %>;\n          var enable_rationale = <%= enable_rationale %>;\n          var short_circuit = <%= short_circuit %>;\n\n        <% if (!test_id.nil? && test_id.class==Moped::BSON::ObjectId) %>\n          var test_id = new ObjectId(\"<%= test_id %>\");\n        <% else %>\n          var test_id = null;\n        <% end %>\n\n          hqmfjs = {}\n          <%= init_js_frameworks %>\n\n          hqmfjs.effective_date = effective_date;\n          hqmfjs.test_id = test_id;\n      \n          \n        var patient_api = new hQuery.Patient(patient);\n\n        \n        // #########################\n        // ##### DATA ELEMENTS #####\n        // #########################\n\n        hqmfjs.nqf_id = '0480';\n        hqmfjs.hqmf_id = '40280381-4DE7-DB4D-014D-E8C552E9025F';\n        hqmfjs.sub_id = 'a';\n        if (typeof(test_id) == 'undefined') hqmfjs.test_id = null;\n\n        OidDictionary = {'2.16.840.1.113883.3.117.1.7.1.25':{'SNOMED-CT':['169826009']},'2.16.840.1.113883.3.117.1.7.1.26':{'ICD-10-CM':['Z38.00','Z38.01'],'ICD-9-CM':['V30.00','V30.01']},'2.16.840.1.113883.3.117.1.7.1.35':{'ICD-9-CM':['271.1'],'ICD-10-CM':['E74.20','E74.21','E74.29'],'SNOMED-CT':['190745006']},'2.16.840.1.113883.3.117.1.7.1.75':{'SNOMED-CT':['405269005']},'2.16.840.1.113883.3.117.1.7.1.87':{'SNOMED-CT':['306701001','434781000124105','306703003']},'2.16.840.1.113762.1.4.1045.47':{'SNOMED-CT':['412726003']},'2.16.840.1.113883.3.117.1.7.1.38':{'ICD-10-PCS':['3E0536Z','3E0436Z','3E0636Z','3E0336Z'],'SNOMED-CT':['183027000','230115000','225372007','230114001','25156005','304055007'],'ICD-9-PCS':['99.15']},'2.16.840.1.113883.3.117.1.7.1.30':{'SNOMED-CT':['226789007']},'2.16.840.1.113883.3.117.1.7.1.27':{'SNOMED-CT':['346712003','284458003','67079006','412414007','284459006','412413001','284461002','11713004','386127005']},'2.16.840.1.113762.1.4.1045.29':{'SNOMED-CT':['169643005']},'2.16.840.1.113883.3.117.1.7.1.309':{'SNOMED-CT':['371828006']},'2.16.840.1.113762.1.4.1':{'AdministrativeGender':['M','F']},'2.16.840.1.114222.4.11.836':{'CDC Race':['2076-8','1002-5','2131-1','2106-3','2028-9','2054-5']},'2.16.840.1.114222.4.11.837':{'CDC Race':['2135-2','2186-5']},'2.16.840.1.114222.4.11.3591':{'Source of Payment Typology':['521','84','6','331','3119','953','3222','512','349','37','41','523','3116','312','3113','5','32126','212','3115','3211','54','112','611','311','333','21','122','39','822','332','32122','82','73','322','32125','3711','121','389','3','511','342','36','3712','59','3221','379','62','43','3223','123','119','3212','32121','52','81','55','34','69','8','821','98','3112','519','3114','79','3811','32123','25','38','613','35','2','94','85','99','91','3123','321','3229','3813','83','3713','24','951','213','522','129','61','1','3122','64','612','334','529','22','33','31','72','219','619','3812','92','211','29','4','51','313','63','341','9','514','3819','513','515','95','44','42','53','823','96','113','93','343','3121','362','89','959','11','32','32124','381','111','23','19','954','12','372','9999','71','361','7','382','369','3111','371']},'2.16.840.1.113883.3.666.5.307':{'SNOMED-CT':['8715000','183452005','32485007']},'2.16.840.1.113883.3.117.1.7.1.70':{'SNOMED-CT':['3950001']}};\n        \n        // Measure variables\nvar MeasurePeriod = {\n  \"low\": new TS(\"201201010000\", true),\n  \"high\": new TS(\"201212312359\", true)\n}\nhqmfjs.MeasurePeriod = function(patient) {\n  return [new hQuery.CodedEntry(\n    {\n      \"start_time\": MeasurePeriod.low.asDate().getTime()/1000,\n      \"end_time\": MeasurePeriod.high.asDate().getTime()/1000,\n      \"codes\": {}\n    }\n  )];\n}\nif (typeof effective_date === 'number') {\n  MeasurePeriod.high.date = new Date(1000*effective_date);\n  // add one minute before pulling off the year.  This turns 12-31-2012 23:59 into 1-1-2013 00:00 => 1-1-2012 00:00\n  MeasurePeriod.low.date = new Date(1000*(effective_date+60));\n  MeasurePeriod.low.date.setFullYear(MeasurePeriod.low.date.getFullYear()-1);\n}\n\n// Data critera\nhqmfjs.GROUP_variable_CHILDREN_12 = function(patient, initialSpecificContext) {\n  var events = UNION(\n    hqmfjs.GROUP_satisfiesAll_CHILDREN_10(patient, initialSpecificContext)\n  );\n  // record the result of the source of the variable to the rationale\n  if(Logger.enable_rationale) Logger.record('GROUP_variable_CHILDREN_12',events);\n  events.specific_occurrence = 'GROUP_variable_CHILDREN_12';\n\n  events.specificContext=new hqmf.SpecificOccurrence(Row.buildForDataCriteria(events.specific_occurrence, events))\n  return events;\n}\n\nhqmfjs.PatientCharacteristicSexOncAdministrativeSex = function(patient, initialSpecificContext) {\n  var value = patient.gender() || null;\n  matching = matchingValue(value, new CD(\"M\", \"Administrative Sex\"));\n  matching.specificContext=hqmf.SpecificsManager.identity();\n  return matching;\n}\n\nhqmfjs.PatientCharacteristicRaceRace = function(patient, initialSpecificContext) {\n  var value = patient.race() || null;\n  matching = new Boolean(value.includedIn({\"CDC Race\":[\"2076-8\",\"1002-5\",\"2131-1\",\"2106-3\",\"2028-9\",\"2054-5\"]}));\n  matching.specificContext=hqmf.SpecificsManager.identity();\n  return matching;\n}\n\nhqmfjs.PatientCharacteristicEthnicityEthnicity = function(patient, initialSpecificContext) {\n  var value = patient.ethnicity() || null;\n  matching = matchingValue(value, null);\n  matching.specificContext=hqmf.SpecificsManager.identity();\n  return matching;\n}\n\nhqmfjs.PatientCharacteristicPayerPayer = function(patient, initialSpecificContext) {\n  var value = patient.payer() || null;\n  matching = matchingValue(value, null);\n  matching.specificContext=hqmf.SpecificsManager.identity();\n  return matching;\n}\n\nhqmfjs.EncounterPerformedEncounterInpatient_precondition_2 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"encounters\", \"statuses\": [\"performed\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.666.5.307\"};\n  var events = patient.getEvents(eventCriteria);\n  events = filterEventsByField(events, \"lengthOfStay\", new IVL_PQ(null, new PQ(120, \"d\", true)));\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.EncounterPerformedEncounterInpatient_precondition_3 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"encounters\", \"statuses\": [\"performed\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.666.5.307\"};\n  var events = patient.getEvents(eventCriteria);\n  if (events.length > 0 || !Logger.short_circuit) events = EDU(events, hqmfjs.MeasurePeriod(patient));\n  if (events.length == 0) events.specificContext=hqmf.SpecificsManager.empty();\n  return events;\n}\n\nhqmfjs.GROUP_satisfiesAll_CHILDREN_4 = function(patient, initialSpecificContext) {\n  var events = INTERSECT(\n    hqmfjs.EncounterPerformedEncounterInpatient_precondition_2(patient, initialSpecificContext),\n    hqmfjs.EncounterPerformedEncounterInpatient_precondition_3(patient, initialSpecificContext)\n  );\n\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.GROUP_variable_CHILDREN_6 = function(patient, initialSpecificContext) {\n  var events = UNION(\n    hqmfjs.GROUP_satisfiesAll_CHILDREN_4(patient, initialSpecificContext)\n  );\n  // record the result of the source of the variable to the rationale\n  if(Logger.enable_rationale) Logger.record('GROUP_variable_CHILDREN_6',events);\n\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.EncounterPerformedEncounterInpatient_precondition_8 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"encounters\", \"statuses\": [\"performed\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.666.5.307\"};\n  var events = patient.getEvents(eventCriteria);\n  events = filterEventsByField(events, \"lengthOfStay\", new IVL_PQ(null, new PQ(120, \"d\", true)));\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.EncounterPerformedEncounterInpatient_precondition_9 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"encounters\", \"statuses\": [\"performed\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.666.5.307\"};\n  var events = patient.getEvents(eventCriteria);\n  if (events.length > 0 || !Logger.short_circuit) events = EDU(events, hqmfjs.MeasurePeriod(patient));\n  if (events.length == 0) events.specificContext=hqmf.SpecificsManager.empty();\n  return events;\n}\n\nhqmfjs.GROUP_satisfiesAll_CHILDREN_10 = function(patient, initialSpecificContext) {\n  var events = INTERSECT(\n    hqmfjs.EncounterPerformedEncounterInpatient_precondition_8(patient, initialSpecificContext),\n    hqmfjs.EncounterPerformedEncounterInpatient_precondition_9(patient, initialSpecificContext)\n  );\n\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.GROUP_variable_CHILDREN_12 = function(patient, initialSpecificContext) {\n  var events = UNION(\n    hqmfjs.GROUP_satisfiesAll_CHILDREN_10(patient, initialSpecificContext)\n  );\n  // record the result of the source of the variable to the rationale\n  if(Logger.enable_rationale) Logger.record('GROUP_variable_CHILDREN_12',events);\n  events.specific_occurrence = 'GROUP_variable_CHILDREN_12';\n\n  events.specificContext=new hqmf.SpecificOccurrence(Row.buildForDataCriteria(events.specific_occurrence, events))\n  return events;\n}\n\nhqmfjs.DiagnosisActiveSingleLiveBirth_precondition_16 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"allProblems\", \"statuses\": [\"active\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.117.1.7.1.25\"};\n  var events = patient.getEvents(eventCriteria);\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.DiagnosisActiveSingleLiveBornNewbornBornInHospital_precondition_17 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"allProblems\", \"statuses\": [\"active\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.117.1.7.1.26\"};\n  var events = patient.getEvents(eventCriteria);\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.GROUP_UNION_CHILDREN_18 = function(patient, initialSpecificContext) {\n  var events = UNION(\n    hqmfjs.DiagnosisActiveSingleLiveBirth_precondition_16(patient, initialSpecificContext),\n    hqmfjs.DiagnosisActiveSingleLiveBornNewbornBornInHospital_precondition_17(patient, initialSpecificContext)\n  );\n\n  if (events.length > 0 || !Logger.short_circuit) events = SDU(events, hqmfjs.GROUP_variable_CHILDREN_12(patient));\n  if (events.length == 0) events.specificContext=hqmf.SpecificsManager.empty();\n  return events;\n}\n\nhqmfjs.DiagnosisActiveGalactosemia_precondition_19 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"allProblems\", \"statuses\": [\"active\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.117.1.7.1.35\"};\n  var events = patient.getEvents(eventCriteria);\n  if (events.length > 0 || !Logger.short_circuit) events = SDU(events, hqmfjs.GROUP_variable_CHILDREN_12(patient));\n  if (events.length == 0) events.specificContext=hqmf.SpecificsManager.empty();\n  return events;\n}\n\nhqmfjs.EncounterPerformedEncounterInpatient_precondition_23 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"encounters\", \"statuses\": [\"performed\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.666.5.307\"};\n  var events = patient.getEvents(eventCriteria);\n  events = filterEventsByField(events, \"facility\", new CodeList(getCodes(\"2.16.840.1.113883.3.117.1.7.1.75\")));\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.EncounterPerformedEncounterInpatient_precondition_24 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"encounters\", \"statuses\": [\"performed\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.666.5.307\"};\n  var events = patient.getEvents(eventCriteria);\n  events = filterEventsByField(events, \"dischargeDisposition\", new CodeList(getCodes(\"2.16.840.1.113883.3.117.1.7.1.309\")));\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.EncounterPerformedEncounterInpatient_precondition_25 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"encounters\", \"statuses\": [\"performed\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.666.5.307\"};\n  var events = patient.getEvents(eventCriteria);\n  events = filterEventsByField(events, \"dischargeDisposition\", new CodeList(getCodes(\"2.16.840.1.113883.3.117.1.7.1.87\")));\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.GROUP_satisfiesAny_CHILDREN_26 = function(patient, initialSpecificContext) {\n  var events = UNION(\n    hqmfjs.EncounterPerformedEncounterInpatient_precondition_23(patient, initialSpecificContext),\n    hqmfjs.EncounterPerformedEncounterInpatient_precondition_24(patient, initialSpecificContext),\n    hqmfjs.EncounterPerformedEncounterInpatient_precondition_25(patient, initialSpecificContext)\n  );\n\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.GROUP_INTERSECT_CHILDREN_27 = function(patient, initialSpecificContext) {\n  var events = INTERSECT(\n    hqmfjs.GROUP_variable_CHILDREN_12(patient, initialSpecificContext),\n    hqmfjs.GROUP_satisfiesAny_CHILDREN_26(patient, initialSpecificContext)\n  );\n\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.PhysicalExamPerformedEstimatedGestationalAgeAtBirth_precondition_28 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"procedureResults\", \"statuses\": [\"performed\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113762.1.4.1045.47\"};\n  var events = patient.getEvents(eventCriteria);\n  events = filterEventsByValue(events, new IVL_PQ(new PQ(37, \"wk\", true), null));\n  if (events.length > 0 || !Logger.short_circuit) events = SDU(events, hqmfjs.GROUP_variable_CHILDREN_12(patient));\n  if (events.length == 0) events.specificContext=hqmf.SpecificsManager.empty();\n  return events;\n}\n\nhqmfjs.ProcedurePerformedParenteralNutrition_precondition_29 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"allProcedures\", \"statuses\": [\"performed\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.117.1.7.1.38\"};\n  var events = patient.getEvents(eventCriteria);\n  if (events.length > 0 || !Logger.short_circuit) events = SDU(events, hqmfjs.GROUP_variable_CHILDREN_12(patient));\n  if (events.length == 0) events.specificContext=hqmf.SpecificsManager.empty();\n  return events;\n}\n\nhqmfjs.SubstanceAdministeredBreastMilk_precondition_30 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"allMedications\", \"statuses\": [\"administered\"], \"includeEventsWithoutStatus\": false, \"valueSetId\": \"2.16.840.1.113883.3.117.1.7.1.30\"};\n  var events = patient.getEvents(eventCriteria);\n  if (events.length > 0 || !Logger.short_circuit) events = SDU(events, hqmfjs.GROUP_variable_CHILDREN_12(patient));\n  if (events.length == 0) events.specificContext=hqmf.SpecificsManager.empty();\n  return events;\n}\n\nhqmfjs.SubstanceAdministeredDietaryIntakeOtherThanBreastMilk_precondition_31 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"allMedications\", \"statuses\": [\"administered\"], \"includeEventsWithoutStatus\": false, \"valueSetId\": \"2.16.840.1.113883.3.117.1.7.1.27\"};\n  var events = patient.getEvents(eventCriteria);\n  if (events.length > 0 || !Logger.short_circuit) events = SDU(events, hqmfjs.GROUP_variable_CHILDREN_12(patient));\n  if (events.length == 0) events.specificContext=hqmf.SpecificsManager.empty();\n  return events;\n}\n\nhqmfjs.PatientCharacteristicBirthdateBirthdate = function(patient, initialSpecificContext) {\n  var value = patient.birthtime() || null;\n  var events = value ? [value] : [];\n  events.specificContext=events.specificContext||hqmf.SpecificsManager.identity();\n  return events;\n}\n\nhqmfjs.CommunicationFromPatientToProviderFeedingIntentionBreast_precondition_32 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"procedures\", \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113762.1.4.1045.29\"};\n  var events = patient.getEvents(eventCriteria);\n  if (events.length > 0 || !Logger.short_circuit) events = SAS(events, hqmfjs.PatientCharacteristicBirthdateBirthdate(patient), new IVL_PQ(null, new PQ(1, \"h\", true)));\n  if (events.length == 0) events.specificContext=hqmf.SpecificsManager.empty();\n  return events;\n}\n\n\n\n        // #########################\n        // ##### MEASURE LOGIC #####\n        // #########################\n        \n        hqmfjs.initializeSpecifics = function(patient_api, hqmfjs) { hqmf.SpecificsManager.initialize(patient_api,hqmfjs,{\"id\":\"GROUP_variable_CHILDREN_12\",\"type\":\"OCCURRENCE_A_OF_ENCOUNTERINPATIENT\",\"function\":\"GROUP_variable_CHILDREN_12\"}) }\n\n        // INITIAL PATIENT POPULATION\n        hqmfjs.IPP = function(patient, initialSpecificContext) {\n  population_criteria_fn = allTrue('IPP', patient, initialSpecificContext,\n    allTrue('33', patient, initialSpecificContext, hqmfjs.GROUP_variable_CHILDREN_12, hqmfjs.PhysicalExamPerformedEstimatedGestationalAgeAtBirth_precondition_28, hqmfjs.GROUP_UNION_CHILDREN_18,\n      atLeastOneTrue('37', patient, initialSpecificContext,\n        allFalse('40', patient, initialSpecificContext, hqmfjs.DiagnosisActiveGalactosemia_precondition_19, hqmfjs.ProcedurePerformedParenteralNutrition_precondition_29\n        )\n      )\n    )\n  );\n  if (typeof(population_criteria_fn) == 'function') {\n  \treturn population_criteria_fn();\n  } else {\n  \treturn population_criteria_fn;\n  }\n};\n\n\n        // STRATIFICATION\n        hqmfjs.STRAT=null;\n        // DENOMINATOR\n        hqmfjs.DENOM = function(patient) { return new Boolean(true); }\n        // NUMERATOR\n        hqmfjs.NUMER = function(patient, initialSpecificContext) {\n  population_criteria_fn = allTrue('NUMER', patient, initialSpecificContext,\n    allTrue('44', patient, initialSpecificContext, hqmfjs.SubstanceAdministeredBreastMilk_precondition_30,\n      atLeastOneTrue('46', patient, initialSpecificContext,\n        allFalse('48', patient, initialSpecificContext, hqmfjs.SubstanceAdministeredDietaryIntakeOtherThanBreastMilk_precondition_31\n        )\n      )\n    )\n  );\n  if (typeof(population_criteria_fn) == 'function') {\n  \treturn population_criteria_fn();\n  } else {\n  \treturn population_criteria_fn;\n  }\n};\n\n\n        hqmfjs.DENEX = function(patient, initialSpecificContext) {\n  population_criteria_fn = atLeastOneTrue('DENEX', patient, initialSpecificContext,\n    atLeastOneTrue('42', patient, initialSpecificContext, hqmfjs.GROUP_INTERSECT_CHILDREN_27\n    )\n  );\n  if (typeof(population_criteria_fn) == 'function') {\n  \treturn population_criteria_fn();\n  } else {\n  \treturn population_criteria_fn;\n  }\n};\n\n\n        hqmfjs.DENEXCEP = function(patient) { return new Boolean(false); }\n        // CV\n        hqmfjs.MSRPOPL = function(patient) { return new Boolean(false); }\n        hqmfjs.OBSERV = function(patient) { return new Boolean(false); }\n        \n        \n        var occurrenceId = [\"GROUP_variable_CHILDREN_12\"];\n\n        hqmfjs.initializeSpecifics(patient_api, hqmfjs)\n        \n        var population = function() {\n          return executeIfAvailable(hqmfjs.IPP, patient_api);\n        }\n        var stratification = null;\n        if (hqmfjs.STRAT) {\n          stratification = function() {\n            return hqmf.SpecificsManager.setIfNull(executeIfAvailable(hqmfjs.STRAT, patient_api));\n          }\n        }\n        var denominator = function() {\n          return executeIfAvailable(hqmfjs.DENOM, patient_api);\n        }\n        var numerator = function() {\n          return executeIfAvailable(hqmfjs.NUMER, patient_api);\n        }\n        var exclusion = function() {\n          return executeIfAvailable(hqmfjs.DENEX, patient_api);\n        }\n        var denexcep = function() {\n          return executeIfAvailable(hqmfjs.DENEXCEP, patient_api);\n        }\n        var msrpopl = function(specific_context) {\n          if (specific_context){\n            \n          var observFunc = hqmfjs.MSRPOPL\n          if (typeof(observFunc)==='function')\n            return observFunc(patient_api, specific_context);\n          else\n            return [];\n          }\n          else {\n            return executeIfAvailable(hqmfjs.MSRPOPL, patient_api);\n          }\n        }\n        var observ = function(specific_context) {\n          \n          var observFunc = hqmfjs.OBSERV\n          if (typeof(observFunc)==='function')\n            return observFunc(patient_api, specific_context);\n          else\n            return [];\n        }\n        \n        var executeIfAvailable = function(optionalFunction, patient_api) {\n          if (typeof(optionalFunction)==='function') {\n            result = optionalFunction(patient_api);\n            \n            return result;\n          } else {\n            return false;\n          }\n        }\n\n        \n        if (typeof Logger != 'undefined') {\n          // clear out logger\n          Logger.logger = [];\n          Logger.rationale={};\n          if (typeof short_circuit == 'undefined') short_circuit = true;\n        \n          // turn on logging if it is enabled\n          if (enable_logging || enable_rationale) {\n            injectLogger(hqmfjs, enable_logging, enable_rationale, short_circuit);\n          } else {\n            Logger.enable_rationale = false;\n          }\n        }\n\n        try {\n          map(patient, population, denominator, numerator, exclusion, denexcep, msrpopl, observ, occurrenceId,false,stratification);\n        } catch(err) {\n          print(err.stack);\n          throw err;\n        }\n\n        \n        };\n        ",
  "continuous_variable": false,
  "episode_of_care": true,
  "hqmf_document": {
    "id": "0480",
    "hqmf_id": "40280381-4DE7-DB4D-014D-E8C552E9025F",
    "hqmf_set_id": "7D374C6A-3821-4333-A1BC-4531005D77B8",
    "hqmf_version_number": 4,
    "title": "Exclusive Breast Milk Feeding",
    "description": "PC-05 Exclusive breast milk feeding during the newborn's entire hospitalization.\n\nPC-05a Exclusive breast milk feeding during the newborn's entire hospitalization considering mother's choice.",
    "cms_id": "CMS9v4",
    "population_criteria": {
      "IPP": {
        "conjunction?": true,
        "type": "IPP",
        "title": "Initial Patient Population",
        "hqmf_id": "A487C868-83A9-4598-9AEC-B50CB6B59D45",
        "preconditions": [
          {
            "id": 33,
            "preconditions": [
              {
                "id": 11,
                "reference": "GROUP_variable_CHILDREN_12"
              },
              {
                "id": 28,
                "reference": "PhysicalExamPerformedEstimatedGestationalAgeAtBirth_precondition_28"
              },
              {
                "id": 14,
                "reference": "GROUP_UNION_CHILDREN_18"
              },
              {
                "id": 37,
                "preconditions": [
                  {
                    "id": 40,
                    "preconditions": [
                      {
                        "id": 19,
                        "reference": "DiagnosisActiveGalactosemia_precondition_19"
                      },
                      {
                        "id": 29,
                        "reference": "ProcedurePerformedParenteralNutrition_precondition_29"
                      }
                    ],
                    "conjunction_code": "atLeastOneTrue",
                    "negation": true
                  }
                ],
                "conjunction_code": "atLeastOneTrue"
              }
            ],
            "conjunction_code": "allTrue"
          }
        ]
      },
      "DENOM": {
        "conjunction?": true,
        "type": "DENOM",
        "title": "Denominator",
        "hqmf_id": "C688C9BD-1A24-4EAB-A64D-BBC366DCFC6C"
      },
      "DENEX": {
        "conjunction?": true,
        "type": "DENEX",
        "title": "Denominator Exclusion",
        "hqmf_id": "F1990F1F-63CF-4027-B09A-6323D80C68F3",
        "preconditions": [
          {
            "id": 42,
            "preconditions": [
              {
                "id": 20,
                "reference": "GROUP_INTERSECT_CHILDREN_27"
              }
            ],
            "conjunction_code": "atLeastOneTrue"
          }
        ]
      },
      "NUMER": {
        "conjunction?": true,
        "type": "NUMER",
        "title": "Numerator",
        "hqmf_id": "6DD7F768-F375-4105-8EBA-8FBCDA3DF982",
        "preconditions": [
          {
            "id": 44,
            "preconditions": [
              {
                "id": 30,
                "reference": "SubstanceAdministeredBreastMilk_precondition_30"
              },
              {
                "id": 46,
                "preconditions": [
                  {
                    "id": 48,
                    "preconditions": [
                      {
                        "id": 31,
                        "reference": "SubstanceAdministeredDietaryIntakeOtherThanBreastMilk_precondition_31"
                      }
                    ],
                    "conjunction_code": "atLeastOneTrue",
                    "negation": true
                  }
                ],
                "conjunction_code": "atLeastOneTrue"
              }
            ],
            "conjunction_code": "allTrue"
          }
        ]
      },
      "IPP_1": {
        "conjunction?": true,
        "type": "IPP",
        "title": "Initial Patient Population",
        "hqmf_id": "257D53B1-89B9-493C-BD38-E9BD5FD13BC1",
        "preconditions": [
          {
            "id": 51,
            "preconditions": [
              {
                "id": 11,
                "reference": "GROUP_variable_CHILDREN_12"
              },
              {
                "id": 28,
                "reference": "PhysicalExamPerformedEstimatedGestationalAgeAtBirth_precondition_28"
              },
              {
                "id": 14,
                "reference": "GROUP_UNION_CHILDREN_18"
              },
              {
                "id": 55,
                "preconditions": [
                  {
                    "id": 58,
                    "preconditions": [
                      {
                        "id": 19,
                        "reference": "DiagnosisActiveGalactosemia_precondition_19"
                      },
                      {
                        "id": 29,
                        "reference": "ProcedurePerformedParenteralNutrition_precondition_29"
                      }
                    ],
                    "conjunction_code": "atLeastOneTrue",
                    "negation": true
                  }
                ],
                "conjunction_code": "atLeastOneTrue"
              }
            ],
            "conjunction_code": "allTrue"
          }
        ]
      },
      "DENOM_1": {
        "conjunction?": true,
        "type": "DENOM",
        "title": "Denominator",
        "hqmf_id": "19719CBA-9F82-4647-843C-D0FC1BDA8EEE"
      },
      "DENEX_1": {
        "conjunction?": true,
        "type": "DENEX",
        "title": "Denominator Exclusion",
        "hqmf_id": "E6523BDD-D4FA-4849-9217-F7A3668B1746",
        "preconditions": [
          {
            "id": 60,
            "preconditions": [
              {
                "id": 20,
                "reference": "GROUP_INTERSECT_CHILDREN_27"
              },
              {
                "id": 62,
                "preconditions": [
                  {
                    "id": 64,
                    "preconditions": [
                      {
                        "id": 32,
                        "reference": "CommunicationFromPatientToProviderFeedingIntentionBreast_precondition_32"
                      }
                    ],
                    "conjunction_code": "allTrue",
                    "negation": true
                  }
                ],
                "conjunction_code": "allTrue"
              }
            ],
            "conjunction_code": "atLeastOneTrue"
          }
        ]
      },
      "NUMER_1": {
        "conjunction?": true,
        "type": "NUMER",
        "title": "Numerator",
        "hqmf_id": "AB070153-6811-452C-8630-0D1A28B85FE6",
        "preconditions": [
          {
            "id": 65,
            "preconditions": [
              {
                "id": 30,
                "reference": "SubstanceAdministeredBreastMilk_precondition_30"
              },
              {
                "id": 67,
                "preconditions": [
                  {
                    "id": 69,
                    "preconditions": [
                      {
                        "id": 31,
                        "reference": "SubstanceAdministeredDietaryIntakeOtherThanBreastMilk_precondition_31"
                      }
                    ],
                    "conjunction_code": "atLeastOneTrue",
                    "negation": true
                  }
                ],
                "conjunction_code": "atLeastOneTrue"
              }
            ],
            "conjunction_code": "allTrue"
          }
        ]
      }
    },
    "data_criteria": {
      "EncounterOrderDecisionToAdmitToHospitalInpatient": {
        "title": "Decision to Admit to Hospital Inpatient",
        "description": "Encounter, Order: Decision to Admit to Hospital Inpatient",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.295",
        "type": "encounters",
        "definition": "encounter",
        "status": "ordered",
        "hard_status": true,
        "negation": false,
        "source_data_criteria": "EncounterOrderDecisionToAdmitToHospitalInpatient",
        "variable": false
      },
      "PatientCharacteristicSexOncAdministrativeSex": {
        "title": "ONC Administrative Sex",
        "description": "Patient Characteristic Sex: ONC Administrative Sex",
        "code_list_id": "2.16.840.1.113762.1.4.1",
        "property": "gender",
        "type": "characteristic",
        "definition": "patient_characteristic_gender",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicSexOncAdministrativeSex",
        "variable": false,
        "value": {
          "type": "CD",
          "system": "Administrative Sex",
          "code": "M"
        }
      },
      "PatientCharacteristicRaceRace": {
        "title": "Race",
        "description": "Patient Characteristic Race: Race",
        "code_list_id": "2.16.840.1.114222.4.11.836",
        "property": "race",
        "type": "characteristic",
        "definition": "patient_characteristic_race",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicRaceRace",
        "variable": false,
        "inline_code_list": {
          "CDC Race": [
            "2076-8",
            "1002-5",
            "2131-1",
            "2106-3",
            "2028-9",
            "2054-5"
          ]
        }
      },
      "PatientCharacteristicEthnicityEthnicity": {
        "title": "Ethnicity",
        "description": "Patient Characteristic Ethnicity: Ethnicity",
        "code_list_id": "2.16.840.1.114222.4.11.837",
        "property": "ethnicity",
        "type": "characteristic",
        "definition": "patient_characteristic_ethnicity",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicEthnicityEthnicity",
        "variable": false,
        "inline_code_list": {
          "CDC Race": [
            "2135-2",
            "2186-5"
          ]
        }
      },
      "PatientCharacteristicPayerPayer": {
        "title": "Payer",
        "description": "Patient Characteristic Payer: Payer",
        "code_list_id": "2.16.840.1.114222.4.11.3591",
        "property": "payer",
        "type": "characteristic",
        "definition": "patient_characteristic_payer",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicPayerPayer",
        "variable": false,
        "inline_code_list": {
          "Source of Payment Typology": [
            "521",
            "84",
            "6",
            "331",
            "3119",
            "953",
            "3222",
            "512",
            "349",
            "37",
            "41",
            "523",
            "3116",
            "312",
            "3113",
            "5",
            "32126",
            "212",
            "3115",
            "3211",
            "54",
            "112",
            "611",
            "311",
            "333",
            "21",
            "122",
            "39",
            "822",
            "332",
            "32122",
            "82",
            "73",
            "322",
            "32125",
            "3711",
            "121",
            "389",
            "3",
            "511",
            "342",
            "36",
            "3712",
            "59",
            "3221",
            "379",
            "62",
            "43",
            "3223",
            "123",
            "119",
            "3212",
            "32121",
            "52",
            "81",
            "55",
            "34",
            "69",
            "8",
            "821",
            "98",
            "3112",
            "519",
            "3114",
            "79",
            "3811",
            "32123",
            "25",
            "38",
            "613",
            "35",
            "2",
            "94",
            "85",
            "99",
            "91",
            "3123",
            "321",
            "3229",
            "3813",
            "83",
            "3713",
            "24",
            "951",
            "213",
            "522",
            "129",
            "61",
            "1",
            "3122",
            "64",
            "612",
            "334",
            "529",
            "22",
            "33",
            "31",
            "72",
            "219",
            "619",
            "3812",
            "92",
            "211",
            "29",
            "4",
            "51",
            "313",
            "63",
            "341",
            "9",
            "514",
            "3819",
            "513",
            "515",
            "95",
            "44",
            "42",
            "53",
            "823",
            "96",
            "113",
            "93",
            "343",
            "3121",
            "362",
            "89",
            "959",
            "11",
            "32",
            "32124",
            "381",
            "111",
            "23",
            "19",
            "954",
            "12",
            "372",
            "9999",
            "71",
            "361",
            "7",
            "382",
            "369",
            "3111",
            "371"
          ]
        }
      },
      "EncounterPerformedEncounterInpatient_precondition_2": {
        "title": "Encounter Inpatient",
        "description": "Encounter, Performed: Encounter Inpatient",
        "code_list_id": "2.16.840.1.113883.3.666.5.307",
        "type": "encounters",
        "definition": "encounter",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "EncounterPerformedEncounterInpatient",
        "variable": false,
        "field_values": {
          "LENGTH_OF_STAY": {
            "type": "IVL_PQ",
            "high": {
              "type": "PQ",
              "unit": "d",
              "value": "120",
              "inclusive?": true,
              "derived?": false
            }
          }
        }
      },
      "EncounterPerformedEncounterInpatient_precondition_3": {
        "title": "Encounter Inpatient",
        "description": "Encounter, Performed: Encounter Inpatient",
        "code_list_id": "2.16.840.1.113883.3.666.5.307",
        "type": "encounters",
        "definition": "encounter",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "EncounterPerformedEncounterInpatient",
        "variable": false,
        "temporal_references": [
          {
            "type": "EDU",
            "reference": "MeasurePeriod"
          }
        ]
      },
      "GROUP_satisfiesAll_CHILDREN_4": {
        "title": "GROUP_satisfiesAll_CHILDREN_4",
        "description": "Encounter Inpatient : Encounter, Performed",
        "children_criteria": [
          "EncounterPerformedEncounterInpatient_precondition_2",
          "EncounterPerformedEncounterInpatient_precondition_3"
        ],
        "derivation_operator": "INTERSECT",
        "type": "derived",
        "definition": "satisfies_all",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "GROUP_satisfiesAll_CHILDREN_4",
        "variable": false
      },
      "GROUP_variable_CHILDREN_6": {
        "title": "GROUP_variable_CHILDREN_6",
        "description": "EncounterInpatient",
        "children_criteria": [
          "GROUP_satisfiesAll_CHILDREN_4"
        ],
        "derivation_operator": "UNION",
        "type": "derived",
        "definition": "derived",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "GROUP_variable_CHILDREN_6",
        "variable": true
      },
      "EncounterPerformedEncounterInpatient_precondition_8": {
        "title": "Encounter Inpatient",
        "description": "Encounter, Performed: Encounter Inpatient",
        "code_list_id": "2.16.840.1.113883.3.666.5.307",
        "type": "encounters",
        "definition": "encounter",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "EncounterPerformedEncounterInpatient",
        "variable": false,
        "field_values": {
          "LENGTH_OF_STAY": {
            "type": "IVL_PQ",
            "high": {
              "type": "PQ",
              "unit": "d",
              "value": "120",
              "inclusive?": true,
              "derived?": false
            }
          }
        }
      },
      "EncounterPerformedEncounterInpatient_precondition_9": {
        "title": "Encounter Inpatient",
        "description": "Encounter, Performed: Encounter Inpatient",
        "code_list_id": "2.16.840.1.113883.3.666.5.307",
        "type": "encounters",
        "definition": "encounter",
        "status": "performed",
        "hqmf_oid": "2.16.840.1.113883.3.560.1.79",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "EncounterPerformedEncounterInpatient",
        "variable": false,
        "temporal_references": [
          {
            "type": "EDU",
            "reference": "MeasurePeriod"
          }
        ]
      },
      "GROUP_satisfiesAll_CHILDREN_10": {
        "title": "GROUP_satisfiesAll_CHILDREN_10",
        "description": "Encounter Inpatient : Encounter, Performed",
        "children_criteria": [
          "EncounterPerformedEncounterInpatient_precondition_8",
          "EncounterPerformedEncounterInpatient_precondition_9"
        ],
        "derivation_operator": "INTERSECT",
        "type": "derived",
        "definition": "satisfies_all",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "GROUP_satisfiesAll_CHILDREN_10",
        "variable": false
      },
      "GROUP_variable_CHILDREN_12": {
        "title": "GROUP_variable_CHILDREN_12",
        "description": "Occurrence A of $EncounterInpatient",
        "children_criteria": [
          "GROUP_satisfiesAll_CHILDREN_10"
        ],
        "derivation_operator": "UNION",
        "type": "derived",
        "definition": "derived",
        "hard_status": false,
        "negation": false,
        "specific_occurrence": "A",
        "specific_occurrence_const": "OCCURRENCE_A_OF_ENCOUNTERINPATIENT",
        "source_data_criteria": "GROUP_variable_CHILDREN_12",
        "variable": true
      },
      "DiagnosisActiveSingleLiveBirth_precondition_16": {
        "title": "Single Live Birth",
        "description": "Diagnosis, Active: Single Live Birth",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.25",
        "type": "conditions",
        "definition": "diagnosis",
        "status": "active",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "DiagnosisActiveSingleLiveBirth",
        "variable": false
      },
      "DiagnosisActiveSingleLiveBornNewbornBornInHospital_precondition_17": {
        "title": "Single Live Born Newborn Born in Hospital",
        "description": "Diagnosis, Active: Single Live Born Newborn Born in Hospital",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.26",
        "type": "conditions",
        "definition": "diagnosis",
        "status": "active",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "DiagnosisActiveSingleLiveBornNewbornBornInHospital",
        "variable": false
      },
      "GROUP_UNION_CHILDREN_18": {
        "title": "GROUP_UNION_CHILDREN_18",
        "description": "",
        "children_criteria": [
          "DiagnosisActiveSingleLiveBirth_precondition_16",
          "DiagnosisActiveSingleLiveBornNewbornBornInHospital_precondition_17"
        ],
        "derivation_operator": "UNION",
        "type": "derived",
        "definition": "derived",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "GROUP_UNION_CHILDREN_18",
        "variable": false,
        "temporal_references": [
          {
            "type": "SDU",
            "reference": "GROUP_variable_CHILDREN_12"
          }
        ]
      },
      "DiagnosisActiveGalactosemia_precondition_19": {
        "title": "Galactosemia",
        "description": "Diagnosis, Active: Galactosemia",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.35",
        "type": "conditions",
        "definition": "diagnosis",
        "status": "active",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "DiagnosisActiveGalactosemia",
        "variable": false,
        "temporal_references": [
          {
            "type": "SDU",
            "reference": "GROUP_variable_CHILDREN_12"
          }
        ]
      },
      "EncounterPerformedEncounterInpatient_precondition_23": {
        "title": "Encounter Inpatient",
        "description": "Encounter, Performed: Encounter Inpatient",
        "code_list_id": "2.16.840.1.113883.3.666.5.307",
        "type": "encounters",
        "definition": "encounter",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "EncounterPerformedEncounterInpatient",
        "variable": false,
        "field_values": {
          "FACILITY_LOCATION": {
            "type": "CD",
            "code_list_id": "2.16.840.1.113883.3.117.1.7.1.75",
            "title": "Neonatal Intensive Care Unit (NICU)"
          }
        }
      },
      "EncounterPerformedEncounterInpatient_precondition_24": {
        "title": "Encounter Inpatient",
        "description": "Encounter, Performed: Encounter Inpatient",
        "code_list_id": "2.16.840.1.113883.3.666.5.307",
        "type": "encounters",
        "definition": "encounter",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "EncounterPerformedEncounterInpatient",
        "variable": false,
        "field_values": {
          "DISCHARGE_STATUS": {
            "type": "CD",
            "code_list_id": "2.16.840.1.113883.3.117.1.7.1.309",
            "title": "Patient Expired"
          }
        }
      },
      "EncounterPerformedEncounterInpatient_precondition_25": {
        "title": "Encounter Inpatient",
        "description": "Encounter, Performed: Encounter Inpatient",
        "code_list_id": "2.16.840.1.113883.3.666.5.307",
        "type": "encounters",
        "definition": "encounter",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "EncounterPerformedEncounterInpatient",
        "variable": false,
        "field_values": {
          "DISCHARGE_STATUS": {
            "type": "CD",
            "code_list_id": "2.16.840.1.113883.3.117.1.7.1.87",
            "title": "Discharge To Acute Care Facility"
          }
        }
      },
      "GROUP_satisfiesAny_CHILDREN_26": {
        "title": "GROUP_satisfiesAny_CHILDREN_26",
        "description": "Encounter Inpatient : Encounter, Performed",
        "children_criteria": [
          "EncounterPerformedEncounterInpatient_precondition_23",
          "EncounterPerformedEncounterInpatient_precondition_24",
          "EncounterPerformedEncounterInpatient_precondition_25"
        ],
        "derivation_operator": "UNION",
        "type": "derived",
        "definition": "satisfies_any",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "GROUP_satisfiesAny_CHILDREN_26",
        "variable": false
      },
      "GROUP_INTERSECT_CHILDREN_27": {
        "title": "GROUP_INTERSECT_CHILDREN_27",
        "description": "",
        "children_criteria": [
          "GROUP_variable_CHILDREN_12",
          "GROUP_satisfiesAny_CHILDREN_26"
        ],
        "derivation_operator": "INTERSECT",
        "type": "derived",
        "definition": "derived",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "GROUP_INTERSECT_CHILDREN_27",
        "variable": false
      },
      "PhysicalExamPerformedEstimatedGestationalAgeAtBirth_precondition_28": {
        "title": "Estimated Gestational Age at Birth",
        "description": "Physical Exam, Performed: Estimated Gestational Age at Birth",
        "code_list_id": "2.16.840.1.113762.1.4.1045.47",
        "type": "physical_exams",
        "definition": "physical_exam",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PhysicalExamPerformedEstimatedGestationalAgeAtBirth",
        "variable": false,
        "value": {
          "type": "IVL_PQ",
          "low": {
            "type": "PQ",
            "unit": "wk",
            "value": "37",
            "inclusive?": true,
            "derived?": false
          }
        },
        "temporal_references": [
          {
            "type": "SDU",
            "reference": "GROUP_variable_CHILDREN_12"
          }
        ]
      },
      "ProcedurePerformedParenteralNutrition_precondition_29": {
        "title": "Parenteral Nutrition",
        "description": "Procedure, Performed: Parenteral Nutrition",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.38",
        "type": "procedures",
        "definition": "procedure",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "ProcedurePerformedParenteralNutrition",
        "variable": false,
        "temporal_references": [
          {
            "type": "SDU",
            "reference": "GROUP_variable_CHILDREN_12"
          }
        ]
      },
      "SubstanceAdministeredBreastMilk_precondition_30": {
        "title": "Breast Milk",
        "description": "Substance, Administered: Breast Milk",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.30",
        "type": "substances",
        "definition": "substance",
        "status": "administered",
        "hard_status": true,
        "negation": false,
        "source_data_criteria": "SubstanceAdministeredBreastMilk",
        "variable": false,
        "temporal_references": [
          {
            "type": "SDU",
            "reference": "GROUP_variable_CHILDREN_12"
          }
        ]
      },
      "SubstanceAdministeredDietaryIntakeOtherThanBreastMilk_precondition_31": {
        "title": "Dietary Intake Other than Breast Milk",
        "description": "Substance, Administered: Dietary Intake Other than Breast Milk",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.27",
        "type": "substances",
        "definition": "substance",
        "status": "administered",
        "hard_status": true,
        "negation": false,
        "source_data_criteria": "SubstanceAdministeredDietaryIntakeOtherThanBreastMilk",
        "variable": false,
        "temporal_references": [
          {
            "type": "SDU",
            "reference": "GROUP_variable_CHILDREN_12"
          }
        ]
      },
      "PatientCharacteristicBirthdateBirthdate": {
        "title": "Birthdate",
        "description": "Patient Characteristic Birthdate: Birthdate",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.70",
        "property": "birthtime",
        "type": "characteristic",
        "definition": "patient_characteristic_birthdate",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicBirthdateBirthdate",
        "variable": false,
        "inline_code_list": {
          "SNOMED-CT": [
            "3950001"
          ]
        }
      },
      "CommunicationFromPatientToProviderFeedingIntentionBreast_precondition_32": {
        "title": "Feeding Intention-Breast",
        "description": "Communication: From Patient to Provider: Feeding Intention-Breast",
        "code_list_id": "2.16.840.1.113762.1.4.1045.29",
        "type": "communications",
        "definition": "communication_from_patient_to_provider",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "CommunicationFromPatientToProviderFeedingIntentionBreast",
        "variable": false,
        "temporal_references": [
          {
            "type": "SAS",
            "reference": "PatientCharacteristicBirthdateBirthdate",
            "range": {
              "type": "IVL_PQ",
              "high": {
                "type": "PQ",
                "unit": "h",
                "value": "1",
                "inclusive?": true,
                "derived?": false
              }
            }
          }
        ]
      }
    },
    "source_data_criteria": {
      "PatientCharacteristicBirthdateBirthdate": {
        "title": "Birthdate",
        "description": "Patient Characteristic Birthdate: Birthdate",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.70",
        "property": "birthtime",
        "type": "characteristic",
        "definition": "patient_characteristic_birthdate",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicBirthdateBirthdate",
        "variable": false,
        "inline_code_list": {
          "SNOMED-CT": [
            "3950001"
          ]
        }
      },
      "SubstanceAdministeredBreastMilk": {
        "title": "Breast Milk",
        "description": "Substance, Administered: Breast Milk",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.30",
        "type": "substances",
        "definition": "substance",
        "status": "administered",
        "hard_status": true,
        "negation": false,
        "source_data_criteria": "SubstanceAdministeredBreastMilk",
        "variable": false
      },
      "SubstanceAdministeredDietaryIntakeOtherThanBreastMilk": {
        "title": "Dietary Intake Other than Breast Milk",
        "description": "Substance, Administered: Dietary Intake Other than Breast Milk",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.27",
        "type": "substances",
        "definition": "substance",
        "status": "administered",
        "hard_status": true,
        "negation": false,
        "source_data_criteria": "SubstanceAdministeredDietaryIntakeOtherThanBreastMilk",
        "variable": false
      },
      "EncounterPerformedEncounterInpatient": {
        "title": "Encounter Inpatient",
        "description": "Encounter, Performed: Encounter Inpatient",
        "code_list_id": "2.16.840.1.113883.3.666.5.307",
        "type": "encounters",
        "definition": "encounter",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "EncounterPerformedEncounterInpatient",
        "variable": false
      },
      "PhysicalExamPerformedEstimatedGestationalAgeAtBirth": {
        "title": "Estimated Gestational Age at Birth",
        "description": "Physical Exam, Performed: Estimated Gestational Age at Birth",
        "code_list_id": "2.16.840.1.113762.1.4.1045.47",
        "type": "physical_exams",
        "definition": "physical_exam",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PhysicalExamPerformedEstimatedGestationalAgeAtBirth",
        "variable": false
      },
      "PatientCharacteristicEthnicityEthnicity": {
        "title": "Ethnicity",
        "description": "Patient Characteristic Ethnicity: Ethnicity",
        "code_list_id": "2.16.840.1.114222.4.11.837",
        "property": "ethnicity",
        "type": "characteristic",
        "definition": "patient_characteristic_ethnicity",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicEthnicityEthnicity",
        "variable": false,
        "inline_code_list": {
          "CDC Race": [
            "2135-2",
            "2186-5"
          ]
        }
      },
      "PatientCharacteristicExpiredExpired": {
        "title": "Expired",
        "description": "Patient Characteristic Expired: Expired",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.309",
        "property": "expired",
        "type": "characteristic",
        "definition": "patient_characteristic_expired",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicExpiredExpired",
        "variable": false,
        "inline_code_list": {
          "SNOMED-CT": [
            "371828006"
          ]
        }
      },
      "CommunicationFromPatientToProviderFeedingIntentionBreast": {
        "title": "Feeding Intention-Breast",
        "description": "Communication: From Patient to Provider: Feeding Intention-Breast",
        "code_list_id": "2.16.840.1.113762.1.4.1045.29",
        "type": "communications",
        "definition": "communication_from_patient_to_provider",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "CommunicationFromPatientToProviderFeedingIntentionBreast",
        "variable": false
      },
      "DiagnosisActiveGalactosemia": {
        "title": "Galactosemia",
        "description": "Diagnosis, Active: Galactosemia",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.35",
        "type": "conditions",
        "definition": "diagnosis",
        "status": "active",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "DiagnosisActiveGalactosemia",
        "variable": false
      },
      "PatientCharacteristicSexOncAdministrativeSex": {
        "title": "ONC Administrative Sex",
        "description": "Patient Characteristic Sex: ONC Administrative Sex",
        "code_list_id": "2.16.840.1.113762.1.4.1",
        "property": "gender",
        "type": "characteristic",
        "definition": "patient_characteristic_gender",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicSexOncAdministrativeSex",
        "variable": false,
        "value": {
          "type": "CD",
          "system": "Administrative Sex",
          "code": "M"
        }
      },
      "ProcedurePerformedParenteralNutrition": {
        "title": "Parenteral Nutrition",
        "description": "Procedure, Performed: Parenteral Nutrition",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.38",
        "type": "procedures",
        "definition": "procedure",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "ProcedurePerformedParenteralNutrition",
        "variable": false
      },
      "PatientCharacteristicPayerPayer": {
        "title": "Payer",
        "description": "Patient Characteristic Payer: Payer",
        "code_list_id": "2.16.840.1.114222.4.11.3591",
        "property": "payer",
        "type": "characteristic",
        "definition": "patient_characteristic_payer",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicPayerPayer",
        "variable": false,
        "inline_code_list": {
          "Source of Payment Typology": [
            "521",
            "84",
            "6",
            "331",
            "3119",
            "953",
            "3222",
            "512",
            "349",
            "37",
            "41",
            "523",
            "3116",
            "312",
            "3113",
            "5",
            "32126",
            "212",
            "3115",
            "3211",
            "54",
            "112",
            "611",
            "311",
            "333",
            "21",
            "122",
            "39",
            "822",
            "332",
            "32122",
            "82",
            "73",
            "322",
            "32125",
            "3711",
            "121",
            "389",
            "3",
            "511",
            "342",
            "36",
            "3712",
            "59",
            "3221",
            "379",
            "62",
            "43",
            "3223",
            "123",
            "119",
            "3212",
            "32121",
            "52",
            "81",
            "55",
            "34",
            "69",
            "8",
            "821",
            "98",
            "3112",
            "519",
            "3114",
            "79",
            "3811",
            "32123",
            "25",
            "38",
            "613",
            "35",
            "2",
            "94",
            "85",
            "99",
            "91",
            "3123",
            "321",
            "3229",
            "3813",
            "83",
            "3713",
            "24",
            "951",
            "213",
            "522",
            "129",
            "61",
            "1",
            "3122",
            "64",
            "612",
            "334",
            "529",
            "22",
            "33",
            "31",
            "72",
            "219",
            "619",
            "3812",
            "92",
            "211",
            "29",
            "4",
            "51",
            "313",
            "63",
            "341",
            "9",
            "514",
            "3819",
            "513",
            "515",
            "95",
            "44",
            "42",
            "53",
            "823",
            "96",
            "113",
            "93",
            "343",
            "3121",
            "362",
            "89",
            "959",
            "11",
            "32",
            "32124",
            "381",
            "111",
            "23",
            "19",
            "954",
            "12",
            "372",
            "9999",
            "71",
            "361",
            "7",
            "382",
            "369",
            "3111",
            "371"
          ]
        }
      },
      "PatientCharacteristicRaceRace": {
        "title": "Race",
        "description": "Patient Characteristic Race: Race",
        "code_list_id": "2.16.840.1.114222.4.11.836",
        "property": "race",
        "type": "characteristic",
        "definition": "patient_characteristic_race",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicRaceRace",
        "variable": false,
        "inline_code_list": {
          "CDC Race": [
            "2076-8",
            "1002-5",
            "2131-1",
            "2106-3",
            "2028-9",
            "2054-5"
          ]
        }
      },
      "DiagnosisActiveSingleLiveBirth": {
        "title": "Single Live Birth",
        "description": "Diagnosis, Active: Single Live Birth",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.25",
        "type": "conditions",
        "definition": "diagnosis",
        "status": "active",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "DiagnosisActiveSingleLiveBirth",
        "variable": false
      },
      "DiagnosisActiveSingleLiveBornNewbornBornInHospital": {
        "title": "Single Live Born Newborn Born in Hospital",
        "description": "Diagnosis, Active: Single Live Born Newborn Born in Hospital",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.26",
        "type": "conditions",
        "definition": "diagnosis",
        "status": "active",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "DiagnosisActiveSingleLiveBornNewbornBornInHospital",
        "variable": false
      },
      "GROUP_variable_CHILDREN_6": {
        "title": "GROUP_variable_CHILDREN_6",
        "description": "EncounterInpatient",
        "children_criteria": [
          "GROUP_satisfiesAll_CHILDREN_4"
        ],
        "derivation_operator": "UNION",
        "type": "derived",
        "definition": "derived",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "GROUP_variable_CHILDREN_6",
        "variable": true
      },
      "GROUP_variable_CHILDREN_12": {
        "title": "GROUP_variable_CHILDREN_12",
        "description": "Occurrence A of $EncounterInpatient",
        "children_criteria": [
          "GROUP_satisfiesAll_CHILDREN_10"
        ],
        "derivation_operator": "UNION",
        "type": "derived",
        "definition": "derived",
        "hard_status": false,
        "negation": false,
        "specific_occurrence": "A",
        "specific_occurrence_const": "OCCURRENCE_A_OF_ENCOUNTERINPATIENT",
        "source_data_criteria": "GROUP_variable_CHILDREN_12",
        "variable": true
      }
    },
    "attributes": [
      {
        "id": "NQF_ID_NUMBER",
        "code": "OTH",
        "value": "0480",
        "name": "NQF ID Number"
      },
      {
        "id": "COPYRIGHT",
        "code": "COPY",
        "value": "LOINC(R) is a registered trademark of the Regenstrief Institute.\n \nThis material contains SNOMED Clinical Terms (R) (SNOMED CT[C]) copyright 2004-2014 International Health Terminology Standards Development Organization. All rights reserved.",
        "name": "Copyright"
      },
      {
        "id": "DISCLAIMER",
        "code": "DISC",
        "value": "These performance measures are not clinical guidelines and do not establish a standard of medical care, and have not been tested for all potential applications. The measures and specifications are provided without warranty.",
        "name": "Disclaimer"
      },
      {
        "id": "MEASURE_SCORING",
        "code": "MSRSCORE",
        "value": "Proportion",
        "name": "Measure Scoring"
      },
      {
        "id": "MEASURE_TYPE",
        "code": "MSRTYPE",
        "value": "\n\t\t\tProcess\n\t\t",
        "name": "Measure Type"
      },
      {
        "id": "STRATIFICATION",
        "code": "STRAT",
        "value": "None",
        "name": "Stratification"
      },
      {
        "id": "RISK_ADJUSTMENT",
        "code": "MSRADJ",
        "value": "None",
        "name": "Risk Adjustment"
      },
      {
        "id": "RATE_AGGREGATION",
        "code": "MSRAGG",
        "value": "None",
        "name": "Rate Aggregation"
      },
      {
        "id": "RATIONALE",
        "code": "RAT",
        "value": "Exclusive breast milk feeding for the first 6 months of neonatal life has long been the expressed goal of World Health Organization (WHO), Department of Health and Human Services (DHHS), American Academy of Pediatrics (AAP) and American College of Obstetricians and Gynecologists (ACOG). ACOG has recently reiterated its position (ACOG, 2007). A recent Cochrane review substantiates the benefits (Kramer et al., 2002). Much evidence has now focused on the prenatal and intrapartum period as critical for the success of exclusive (or any) BF (Centers for Disease Control and Prevention [CDC], 2007; Petrova et al., 2007; Shealy et al., 2005; Taveras et al., 2004). Exclusive breast milk feeding rate during birth hospital stay has been calculated by the California Department of Public Health for the last several years using newborn genetic disease testing data. Healthy People 2010 and the CDC have also been active in promoting this goal.",
        "name": "Rationale"
      },
      {
        "id": "CLINICAL_RECOMMENDATION_STATEMENT",
        "code": "CRS",
        "value": "Exclusive breast milk feeding for the first 6 months of neonatal life can result in numerous long-term health benefits for both mother and newborn and is recommended by a number of national and international organizations. Evidence suggests that the prenatal and intrapartum period is critical for the success of exclusive (or any) breast feeding. Therefore, it is recommended that newborns are fed breast milk only from birth to discharge.",
        "name": "Clinical Recommendation Statement"
      },
      {
        "id": "IMPROVEMENT_NOTATION",
        "code": "IDUR",
        "value": "Improvement noted as an increase in the rate",
        "name": "Improvement Notation"
      },
      {
        "id": "REFERENCE",
        "code": "REF",
        "value": "\n\t\t\tAmerican Academy of Pediatrics. (2005). Section on Breastfeeding. Policy Statement:\nBreastfeeding and the Use of Human Milk. Pediatrics.115:496-506.\n\t\t\tAmerican College of Obstetricians and Gynecologists. (Feb. 2007). Committee on Obstetric Practice and Committee on Health Care for Underserved  Women.Breastfeeding: Maternal and Infant Aspects. ACOG Committee Opinion 361.\n\t\t\tCalifornia Department of Public Health. (2006). Genetic Disease Branch. California In-\nHospital Breastfeeding as Indicated on the Newborn Screening Test Form, Statewide, County and Hospital of Occurrence: Available at: http://www.cdph.ca.gov/data/statistics/Pages/BreastfeedingStatistics.aspx.\n\t\t\tCenters for Disease Control and Prevention. (Aug 3, 2007). Breastfeeding trends and\nupdated national health objectives for exclusive breastfeeding--United States birth years 2000-2004. MMWR - Morbidity & Mortality Weekly Report. 56(30):760-3.\n\t\t\tCenters for Disease Control and Prevention. (2014). Division of Nutrition, Physical Activity and Obesity. Breastfeeding Report Card. Available at: http://www.cdc.gov/breastfeeding/pdf/2014breastfeedingreportcard.pdf\n\t\t\tIp, S., Chung, M., Raman, G., et al. (2007). Breastfeeding and maternal and infant health outcomes in developed countries. Rockville, MD: US Department of Health and Human Services. Available at: http://www.ahrq.gov/downloads/pub/evidence/pdf/brfout/brfout.pdf\n\t\t\tKramer, M.S. & Kakuma, R. (2002).Optimal duration of exclusive breastfeeding. [107 refs] Cochrane Database of Systematic Reviews. (1):CD003517.\n\t\t\tPetrova, A., Hegyi, T., & Mehta, R. (2007). Maternal race/ethnicity and one-month exclusive breastfeeding in association with the in-hospital feeding modality. Breastfeeding Medicine. 2(2):92-8.\n\t\t\tShealy, K.R., Li, R., Benton-Davis, S., & Grummer-Strawn, L.M. (2005).The CDC guide to breastfeeding interventions. Atlanta, GA: US Department of Health and Human Services, CDC. Available at: http://www.cdc.gov/breastfeeding/pdf/breastfeeding_interventions.pdf.\n\t\t\tTaveras, E.M., Li, R., Grummer-Strawn, L., Richardson, M., Marshall, R., Rego, V.H.,\nMiroshnik, I., & Lieu, T.A. (2004). Opinions and practices of clinicians associated with\ncontinuation of exclusive breastfeeding. Pediatrics. 113(4):e283-90.\n\t\t\tUS Department of Health and Human Services. (2007). Healthy People 2010 Midcourse Review. Washington, DC: US Department of Health and Human Services. Available at: http://www.healthypeople.gov/2010/data/midcourse/?visit=1.\n\t\t\tWorld Health Organization. (1991). Indicators for assessing breastfeeding practices. Geneva, Switzerland: World Health Organization. Available at: http://whqlibdoc.who.int/hq/1991/WHO_CDD_SER_91.14.pdf?ua=1.\n\t\t",
        "name": "Reference"
      },
      {
        "id": "DEFINITION",
        "code": "DEF",
        "value": "None",
        "name": "Definition"
      },
      {
        "id": "GUIDANCE",
        "code": "GUIDE",
        "value": "A discharge to a designated cancer center or children's hospital should be captured as a discharge to an acute care facility.\n\nThe unit of measurement for this measure is an inpatient episode of care. Each distinct hospitalization should be reported, regardless of whether the same patient is admitted for inpatient care more than once during the measurement period. In addition, the eMeasure logic intends to represent events within or surrounding a single occurrence of an inpatient hospitalization.",
        "name": "Guidance"
      },
      {
        "id": "TRANSMISSION_FORMAT",
        "code": "OTH",
        "value": "TBD",
        "name": "Transmission Format"
      },
      {
        "id": "DENOMINATOR",
        "code": "DENOM",
        "value": "PC-05 Single term newborns discharged from the hospital.\nPC-05a Single term newborns discharged from the hospital.",
        "name": "Denominator"
      },
      {
        "id": "DENOMINATOR_EXCLUSIONS",
        "code": "OTH",
        "value": "PC-05 Newborns who were admitted to the Neonatal Intensive Care Unit (NICU), who were transferred to an acute care facility, or who expired during the hospitalization.\nPC-05a Newborns who were admitted to the Neonatal Intensive Care Unit (NICU), who were transferred to an acute care facility, who expired during the hospitalization, or whose mothers chose not to exclusively breast feed.",
        "name": "Denominator Exclusions"
      },
      {
        "id": "NUMERATOR",
        "code": "NUMER",
        "value": "PC-05 Newborns who were fed breast milk only since birth.\nPC-05a Newborns who were fed breast milk only since birth.",
        "name": "Numerator"
      },
      {
        "id": "NUMERATOR_EXCLUSIONS",
        "code": "OTH",
        "value": "Not applicable",
        "name": "Numerator Exclusions"
      },
      {
        "id": "DENOMINATOR_EXCEPTIONS",
        "code": "DENEXCEP",
        "value": "None",
        "name": "Denominator Exceptions"
      },
      {
        "id": "MEASURE_POPULATION",
        "code": "MSRPOPL",
        "value": "Not applicable",
        "name": "Measure Population"
      },
      {
        "id": "MEASURE_OBSERVATIONS",
        "code": "OTH",
        "value": "Not applicable",
        "name": "Measure Observations"
      },
      {
        "id": "SUPPLEMENTAL_DATA_ELEMENTS",
        "code": "OTH",
        "value": "For every patient evaluated by this measure also identify payer, race, ethnicity and sex.",
        "name": "Supplemental Data Elements"
      }
    ],
    "populations": [
      {"IPP":"IPP","DENOM":"DENOM","DENEX":"DENEX","NUMER":"NUMER","id":"Population1","title":"All"},
      {"IPP":"IPP_1","DENOM":"DENOM_1","DENEX":"DENEX_1","NUMER":"NUMER_1","id":"Population2","title":"Refusals excluded"}
    ],
    "measure_period": {
      "type": "IVL_TS",
      "low": {
        "type": "TS",
        "value": "201201010000",
        "inclusive?": true,
        "derived?": false
      },
      "high": {
        "type": "TS",
        "value": "201212312359",
        "inclusive?": true,
        "derived?": false
      },
      "width": {
        "type": "PQ",
        "unit": "a",
        "value": "1",
        "inclusive?": true,
        "derived?": false
      }
    }
  },
  "sub_id": "a",
  "subtitle": "All",
  "short_subtitle": "All",
  "oids": [
    "2.16.840.1.113883.3.117.1.7.1.25",
    "2.16.840.1.113883.3.117.1.7.1.26",
    "2.16.840.1.113883.3.117.1.7.1.35",
    "2.16.840.1.113883.3.117.1.7.1.75",
    "2.16.840.1.113883.3.117.1.7.1.87",
    "2.16.840.1.113762.1.4.1045.47",
    "2.16.840.1.113883.3.117.1.7.1.38",
    "2.16.840.1.113883.3.117.1.7.1.30",
    "2.16.840.1.113883.3.117.1.7.1.27",
    "2.16.840.1.113762.1.4.1045.29",
    "2.16.840.1.113883.3.117.1.7.1.309",
    "2.16.840.1.113762.1.4.1",
    "2.16.840.1.114222.4.11.836",
    "2.16.840.1.114222.4.11.837",
    "2.16.840.1.114222.4.11.3591",
    "2.16.840.1.113883.3.666.5.307",
    "2.16.840.1.113883.3.117.1.7.1.70"
  ],
  "population_ids": {
    "IPP": "A487C868-83A9-4598-9AEC-B50CB6B59D45",
    "DENOM": "C688C9BD-1A24-4EAB-A64D-BBC366DCFC6C",
    "DENEX": "F1990F1F-63CF-4027-B09A-6323D80C68F3",
    "NUMER": "6DD7F768-F375-4105-8EBA-8FBCDA3DF982"
  }
}
`)
