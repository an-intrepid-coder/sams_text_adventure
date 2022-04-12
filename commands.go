package main

/* For now, Commands will be in the form of "Verb Noun", but this will
   get more flexible and complicated down the road with Modifier particles
   and other grammatical forms for more discerning behavior.  */

var Verbs = MakeEnum([]string{
    "quit",
    "describe",
    "fight",
    "engage",
    "disengage",
})

var Nouns = MakeEnum([]string{
    "self",
    "scene",
})

