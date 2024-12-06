import { AbilityBuilder, Ability } from '@casl/ability'
import themeConfig from './themeConfig'

export const AppAbility = Ability

/**
 * Please define your own Ability rules according to your app requirements.
 * We have just shown Admin and Client rules for demo purpose where
 * admin can manage everything and client can just visit ACL page
 */
const defineRulesFor = (role, permission, permissions) => {
  const { can, rules } = new AbilityBuilder(AppAbility)

  if (permissions?.length) can(['read'], [...permissions])
  if (!themeConfig.acl) {
    can('manage', 'all')

    return rules
  }

  const memberPermissions = [
    'home',
    'challenges',
    'team',
    'team-members',
    'team-settings'
  ]

  switch (role) {
    case 'admin':
      can('manage', 'all')
      break

    case 'member':
      can(['read'], "wallet")
      can(['read'], memberPermissions)
      break

    case 'nowallet-member':
      can(['read'], memberPermissions)
      break

    case 'First-Login':
      can(['read'], "wallet-register")
      break

    default:
      can(['read'], permission)
      break
  }

  return rules
}

export const buildAbilityFor = (role, permission, permissions) => {
  return new AppAbility(defineRulesFor(role, permission, permissions), {
    // https://casl.js.org/v5/en/guide/subject-type-detection
    // @ts-ignore
    detectSubjectType: object => object.type
  })
}

export const defaultACLObj = {
  action: 'manage',
  permission: 'all'
}

export default defineRulesFor
