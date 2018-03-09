<html>
	<head>
		<title>Accelerating SQL Database Operations on a GPU with CUDA</title>
	</head>
	<style>
		body {
			background: #FFFFFF;
			color: #333333;
			text-align: center;
		}
		a {
			color: #000000;
		}
		h1 {
			width: 600px;
			font-size: 34pt;
			font-family: Palatino, Georgia, Arial;
			margin-left: auto;
			margin-right: auto;
			color: #880000;
			text-align: right;
		}
		fieldset {
			background: #FFFFFF;
			width: 600px;
			margin-left: auto;
			margin-right: auto;
			padding: 5px;
			border: 1px solid #CCCCCC;
			text-align: left;
		}
		legend {
			color: #555555;
			font-family: Arial;
			font-weight: bold;
			font-size: 18px;
			font-style: italic;
			padding-right: 10px;
			padding-left: 10px;
		}
	</style>
	<body>
	<script type="text/javascript">

	  var _gaq = _gaq || [];
	    _gaq.push(['_setAccount', 'UA-16432335-1']);
	      _gaq.push(['_trackPageview']);

	        (function() {
		    var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
		        ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
			    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
			      })();

      </script>
		<br/><br/><br/>
		<h1>
			Accelerating SQL Database Operations on a GPU with CUDA
		</h1>
		<fieldset>
			<legend>Abstract</legend>
			Prior work has shown dramatic acceleration for various data-base operations on GPUs, but only using primitives that are not part of conventional database languages such as SQL.  This paper implements a subset of the SQLite virtual machine directly on the GPU, accelerating SQL queries by executing in parallel on GPU hardware.  This dramatically reduces the effort required to achieve GPU acceleration by avoiding the need for database programmers to use new programming languages such as CUDA or modify their programs to use non-SQL libraries.
			<br/><br/>
			This paper focuses on accelerating SELECT queries and describes the considerations in an efficient GPU implementation of the SQLite command processor.  Results on an NVIDIA Tesla C1060 achieve speedups of 20-70x depending on the size of the result set. This work is compared to the alternative query platforms developed recently, such as MapReduce implementations, and future avenues of research in GPU data processing are described.
		</fieldset><br/>
		<fieldset>
			<legend>Papers</legend>
			P. Bakkum and K. Skadron. <a href="bakkum.sql.db.gpu.pdf">Accelerating SQL Database Operations on a GPU with CUDA</a>.
			In <i>GPGPU '10: Proceedings of the Third Workshop on General-Purpose Computation on Graphics
			Processing Units</i>, pages 94-103, New York, NY, USA, 2010. ACM. <br/>
			<a href="bakkum.gpgpu3.pptx">(presentation)</a>

			<br/><br/>

			P.Bakkum and K. Skadron. <a href="bakkum.sql.db.gpu.extended.pdf">Accelerating SQL Database Operations on a GPU with CUDA: Extended Results</a>. University of Virginia Department of Computer Science Technical Report CS-2010-08, May 2010.
		</fieldset><br/>
		<fieldset>
			<legend>Implementation</legend>
			The code associated with this project is available below. The code is intended only for research purposes only and we make no guarantees about performance on your own machine.
			<br/><br/>
			<a href="sphyraena.tar.gz">sphyraena.tar.gz</a> (2526 KB)<br/>
			<a href="sphyraena.zip">sphyraena.zip</a> (2533 KB)
			<br/><br/>
			The code is currently copyright Peter Bakkum, permission is provided to use it for research.
			<br/><br/>
			The code makes use of the SQLite database. Documentation and source code is provided at the <a href="http://sqlite.org">SQLite website</a>. SQLite is in the public domain.
		</fieldset><br/>
		<fieldset>
			<legend>Authors</legend>
			Peter Bakkum is a graduate of the University of Virginia. He is the contact for any issues related to the source code and can be reached at <a href="mailto:pbb7c@virginia.edu">pbb7c@virginia.edu</a>.
			<br/><br/>
			Kevin Skadron is a Professor at the University of Virginia Department of Computer Science. More information can be found on <a href="http://www.cs.virginia.edu/~skadron">his website</a>.
		</fieldset><br/>
		<fieldset>
			<legend>Copyright</legend>
			The content of this webpage and source code linked therein is copyright Peter Bakkum.<br/>
			The content of "Accelerating SQL Database Operations on a GPU with CUDA" is copyright the ACM.<br/><br/>
			Please note that papers linked here represent author preprints.   The
			official, published version must be obtained from the publisher's
			website or the published print copy.  This material is presented here to
			ensure timely dissemination of scholarly and technical work. Copyright
			and all rights therein are retained by authors or by other copyright
			holders. All persons copying this information are expected to adhere to
			the terms and constraints invoked by each document's copyright terms. In
			most cases, these works may not be reposted without the explicit
			permission of the copyright holder. Permission is given to make digital
			or hard copies of all or part of this material without fee for personal
			or classroom use, provided that the copies are not made or distributed
			for profit or commercial advantage, and that copies bear the appropriate
			copyright notice and the full bibliographic citation on the first page.
			 Copyrights for components of this work owned by others must also be
			 honored. To copy otherwise, to republish, to post on servers, to
			 redistribute to lists, etc. requires specific permission and/or a fee.
			 In particular, permission to reprint/republish this material for
			 advertising or promotional purposes or for creating new collective works
			 for resale or redistribution to servers or lists, or to reuse any
			 copyrighted component of this work in other works, must be obtained from
			 the copyright owner.
			<br/><br/>
			 Please note further that any opinions, findings, conclusions, or
			 recommendations expressed in this material are those of the authors and
			 do not necessarily reflect the views of the sponsoring agencies,
			 employers, or publishers.
		</fieldset><br/>
	</body>
</html>
