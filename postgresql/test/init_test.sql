CREATE EXTENSION IF NOT EXISTS pgtap;

SELECT plan( 2 );

SELECT schemas_are(ARRAY[ 'public', 'risky_public', 'risky_private', 'risky_secret' ]);

SELECT tables_are('risky_public',
    ARRAY[ 'action', 'asset', 'attack_chain', 'attack_chain_step', 
    'business', 'capability', 'detection', 'impact', 'liability', 'mitigation',
    'resource', 'threat', 'vulnerability' ]);

SELECT can('risky_public',
    ARRAY[ 'actions','assets','attack_chain_steps','attack_chains','businesses','capabilities','create_action','create_asset','create_attack_chain','create_attack_chain_step','create_business','create_capability','create_detection','create_impact','create_liability','create_mitigation','create_resource','create_threat','create_vulnerability','delete_action','delete_asset','delete_attack_chain','delete_attack_chain_step','delete_business','delete_capability','delete_detection','delete_impact','delete_liability','delete_mitigation','delete_resource','delete_threat','delete_vulnerability','detections','get_action','get_asset','get_attack_chain_step','get_attack_chains','get_business','get_capability','get_detection','get_impact','get_liability','get_mitigation','get_resource','get_threat','get_vulnerability','impacts','liabilities','mitigations','resources','threats','update_action','update_asset','update_attack_chain','update_attack_chain_step','update_business','update_capability','update_detection','update_impact','update_liability','update_mitigation','update_resource','update_threat','update_vulnerability','vulnerabilities']);

--Put single quotes around the array elements 

