package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var resource Resource

	for {
		res, openerErr := opener()
		if openerErr != nil {
			switch openerErr.(type) {
			case TransientError:
				continue
			default:
				return openerErr
			}
		} else {
			resource = res
			break
		}
	}

	if resource != nil {
		defer resource.Close()
	}

	defer func() {
		if r := recover(); r != nil {
			if panicErr, ok := r.(FrobError); ok {
				err = panicErr
				resource.Defrob(panicErr.defrobTag)
			} else {
				err = r.(error)
			}
		}
	}()

	resource.Frob(input)

	return
}
