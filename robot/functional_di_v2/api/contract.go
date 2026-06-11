package api

import "clean-architecture/robot/functional_di_v2/internal/domain"

type Execute func(domain.Cmd)
