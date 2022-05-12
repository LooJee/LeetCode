package first_bad_version

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l := 1
	r := n
	var version int
	for l <= r {
		version = (l + r) / 2
		if isBadVersion(version) {
			if version-1 >= 1 && !isBadVersion(version-1) {
				return version
			}
			r = version - 1
		} else {
			l = version + 1
		}
	}

	return version
}

func isBadVersion(version int) bool {
	return false
}
