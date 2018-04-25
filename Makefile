.PHONY :  get_depend

all: get_depend

get_depend:
	export  $(cat /etc/*release*  | grep =)
	ID=$(ID)
	case ${ID} in 
	gentoo)
		echo 'emerge -avt'
		;;
	debian)
		echo 'apt-get install'
		;;
		*)
		echo 'aaaa'
		;;
	esac
