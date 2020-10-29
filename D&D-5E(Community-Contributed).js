// ==UserScript==
// @name			Roll20 Macro buttons
// @version			0.2
// @description		Adds the roll20 macro calls to the sidebar form the cahracter sheet
// @author			Ben Samsom
// @match			https://app.roll20.net/editor/
// @source          https://github.com/bsamsom/D&D-5E(Community-Contributed).js
// @grant			GM_setValue
// @grant			GM_getValue
// ==/UserScript==
// Changelog: Implemented Saves and ability scores

(function() {
    'use strict';
    var charswitcher = function(event) {
        var e = window.event || event;
        if (e.target.tagName !== 'BUTTON') {
            return;
        }
        var id='';
        var character='';
        if(e.target.hasAttribute('type') && e.target.getAttribute('type') === 'roll') {
            character = e.target;
            while (!character.hasAttribute('data-characterid')) {
                if(!character.parentNode) {
                    return;
                }
                character = character.parentNode;
            }
            id = 'character|' + character.getAttribute('data-characterid');
        } else if(e.target.hasAttribute('class') && e.target.getAttribute('class') === 'btn') {
            var macro = e.target;
            while (!macro.hasAttribute('data-macroid')) {
                if(!macro.parentNode) {
                    return;
                }
                macro = macro.parentNode;
            }
            id = 'character|' + (macro.getAttribute('data-macroid')).split('|')[0];
        }
        if(!id) {
            return;
        }
        var select = document.getElementById('speakingas');
        var x = Object.values(character);
        console.log(x);

        for (var i = 0; i < select.options.length; i++) {
            if (select.options[i].value === id) {
                select.selectedIndex = i;
                return;
            }
        }
    };

    document.getElementsByTagName('body')[0].addEventListener('mousedown', charswitcher);
    document.getElementsByClassName('tokenactions')[0].addEventListener('mousedown', charswitcher);

    var body = document.getElementById("zoomclick");
    var initiative = makeButton("INITIATIVE", "Initiative");
    var death = makeButton("DEATH", "Death_Saving_Throw");
    var saves = saves_table();
    var abilities = ability_table();
    body.appendChild(initiative)
    body.appendChild(saves)
    body.appendChild(death)
    body.appendChild(abilities)



})();

function ability_table() {
    var table = document.createElement('Table');
    var acrobatics     = makeButton("Acrobatics"    , "Acrobatics_Check");
    var medicine       = makeButton("Medicine"      , "Medicine_Check");
    var animalHandling = makeButton("Animal Handling", "AnimalHandling_Check");
    var nature         = makeButton("Nature"        , "Nature_Check");
    var arcana         = makeButton("Arcana"        , "Arcana_Check");
    var perception     = makeButton("Perception"    , "Perception_Check");

    var athletics      = makeButton("Athletics"     , "Athletics_Check");
    var performance    = makeButton("Performance"   , "Performance_Check");
    var deception      = makeButton("Deception"     , "Deception_Check");
    var persuasion     = makeButton("Persuasion"    , "Persuasion_Check");
    var history        = makeButton("History"       , "History_Check");
    var religion       = makeButton("Religion"      , "Religion_Check");

    var insight        = makeButton("Insight"       , "Insight_Check");
    var sleightOfHand  = makeButton("Sleight Of Hand" , "SleightOfHand_check");
    var intimidation   = makeButton("Intimidation"  , "Intimidation_Check");
    var stealth        = makeButton("Stealth"       , "Stealth_check");
    var investigation  = makeButton("Investigation" , "Investigation_Check");
    var survival       = makeButton("Survival"      , "Survival_Check");

    makeRow(acrobatics    , medicine     , table, 0);
    makeRow(animalHandling, nature       , table, 1);
    makeRow(arcana        , perception   , table, 2);
    makeRow(athletics     , performance  , table, 3);
    makeRow(deception     , persuasion   , table, 4);
    makeRow(history       , religion     , table, 5);
    makeRow(insight       , sleightOfHand, table, 6);
    makeRow(intimidation  , stealth      , table, 7);
    makeRow(investigation , survival     , table, 8);
    return table
}

function saves_table() {
    var table = document.createElement('Table');
    var str   = makeButton("STR"  , "Strength_Save");
    var dex   = makeButton("DEX"  , "Dexterity_Save");
    var con   = makeButton("CON"  , "Constitution_Save");
    var int   = makeButton("INT"  , "Intelligence_Save");
    var wis   = makeButton("WIS"  , "Wisdom_Save");
    var cha   = makeButton("CHA"  , "Charisma_Save");
    makeRow(str, dex, table, 0);
    makeRow(con, int, table, 1);
    makeRow(wis, cha, table, 2);
    return table
}

function makeRow(button1, button2, table, pos) {
    var row = table.insertRow(pos);
    var cell1 = row.insertCell(0);
    var cell2 = row.insertCell(1);
    //row.style.width   = "100%"
    //cell1.style.width = "100%"
    //cell2.style.width = "100%"
    cell1.appendChild(button1);
    cell2.appendChild(button2);
    return row;
}

function makeButton(name, stat) {
   var button = document.createElement("Button");
   button.innerHTML = name;
   button.style.width = "100%"
   button.addEventListener ("click", function() {
        $('#textchat-input > textarea').val('%{selected|' + stat + '}');
        $('#textchat-input > button').click();
   });

  return button;
}
