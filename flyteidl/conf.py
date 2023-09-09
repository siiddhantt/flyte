# -*- coding: utf-8 -*-
#
# Configuration file for the Sphinx documentation builder.
#
# This file does only contain a selection of the most common options. For a
# full list see the documentation:
# http://www.sphinx-doc.org/en/stable/config

# -- Path setup --------------------------------------------------------------

# If extensions (or modules to document with autodoc) are in another directory,
# add these directories to sys.path here. If the directory is relative to the
# documentation root, use os.path.abspath to make it absolute, like shown here.
#
import os
import re
import sys

import recommonmark
import sphinx_fontawesome
from recommonmark.transform import AutoStructify


# -- Project information -----------------------------------------------------

project = "Flyte Language Specification"
copyright = "2021, Flyte"
author = "Flyte"

# The full version, including alpha/beta/rc tags
release = re.sub("^v", "", os.popen("git describe").read().strip())
version = release

# -- General configuration ---------------------------------------------------

# Add any Sphinx extension module names here, as strings. They can be
# extensions coming with Sphinx (named 'sphinx.ext.*') or your custom
# ones.
extensions = [
    "sphinx.ext.autodoc",
    "sphinx.ext.autosummary",
    "sphinx.ext.autosectionlabel",
    "sphinx.ext.napoleon",
    "sphinx.ext.todo",
    "sphinx.ext.viewcode",
    "sphinx.ext.doctest",
    "sphinx.ext.intersphinx",
    "sphinx.ext.coverage",
    "sphinx-prompt",
    "sphinx_copybutton",
    "recommonmark",
    "sphinx_markdown_tables",
    "sphinx_fontawesome",
    "sphinx_panels",
]

# build the templated autosummary files
autosummary_generate = True

# autosectionlabel throws warnings if section names are duplicated.
# The following tells autosectionlabel to not throw a warning for
# duplicated section names that are in different documents.
autosectionlabel_prefix_document = True

# Add any paths that contain templates here, relative to this directory.
templates_path = ["_templates"]

# The suffix(es) of source filenames.
# You can specify multiple suffix as a list of string:
#
source_suffix = [".rst"]

# The master toctree document.
master_doc = "index"

# The language for content autogenerated by Sphinx. Refer to documentation
# for a list of supported languages.
#
# This is also used if you do content translation via gettext catalogs.
# Usually you set "language" from the command line for these cases.
language = None

# List of patterns, relative to source directory, that match files and
# directories to ignore when looking for source files.
# This pattern also affects html_static_path and html_extra_path .
exclude_patterns = [
    u"_build",
    "Thumbs.db",
    ".DS_Store",
    "tmp/doc_gen_deps",
    "gen/*/*/*/*/*",
    "CODE_OF_CONDUCT.md",
    "pull_request_template.md",
    "boilerplate",
]

# The name of the Pygments (syntax highlighting) style to use.
pygments_style = "tango"
pygments_dark_style = "native"


# -- Options for HTML output -------------------------------------------------

# The theme to use for HTML and HTML Help pages.  See the documentation for
# a list of builtin themes.
#
html_theme = "furo"
html_title = "Flyte"
html_favicon = "docs/images/flyte_circle_gradient_1_4x4.png"
html_logo = "docs/images/flyte_circle_gradient_1_4x4.png"

html_theme_options = {
    "light_css_variables": {
        "color-brand-primary": "#4300c9",
        "color-brand-content": "#4300c9",
    },
    "dark_css_variables": {
        "color-brand-primary": "#9D68E4",
        "color-brand-content": "#9D68E4",
    },
    # custom flyteorg furo theme options
    "github_repo": "flyteidl",
    "github_username": "flyteorg",
    "github_commit": "master",
    "docs_path": ".",  # path to documentation source
}

html_context = {
    "home_page": "https://docs.flyte.org",
}

# The default sidebars (for documents that don't match any pattern) are
# defined by theme itself.  Builtin themes are using these templates by
# default: ``['localtoc.html', 'relations.html', 'sourcelink.html',
# 'searchbox.html']``.
# html_sidebars = {"**": ["logo-text.html", "globaltoc.html", "localtoc.html", "searchbox.html"]}


# Add any paths that contain custom static files (such as style sheets) here,
# relative to this directory. They are copied after the builtin static files,
# so a file named "default.css" will overwrite the builtin "default.css".
html_static_path = ["docs/_static"]
html_css_files = ["css/fix_toc.css"]

# Custom sidebar templates, must be a dictionary that maps document names
# to template names.
#


# -- Options for HTMLHelp output ---------------------------------------------

# Output file base name for HTML help builder.
htmlhelp_basename = "flyteidldoc"


# -- Options for LaTeX output ------------------------------------------------

latex_elements = {
    # The paper size ('letterpaper' or 'a4paper').
    #
    # 'papersize': 'letterpaper',
    # The font size ('10pt', '11pt' or '12pt').
    #
    # 'pointsize': '10pt',
    # Additional stuff for the LaTeX preamble.
    #
    # 'preamble': '',
    # Latex figure (float) alignment
    #
    # 'figure_align': 'htbp',
}

# Grouping the document tree into LaTeX files. List of tuples
# (source start file, target name, title,
#  author, documentclass [howto, manual, or own class]).
latex_documents = [
    (master_doc, "flyteidl.tex", "flyteidl Documentation", "Flyte", "manual"),
]


# -- Options for manual page output ------------------------------------------

# One entry per manual page. List of tuples
# (source start file, name, description, authors, manual section).
man_pages = [(master_doc, "flyteidl", "flyteidl Documentation", [author], 1)]


# -- Options for Texinfo output ----------------------------------------------

# Grouping the document tree into Texinfo files. List of tuples
# (source start file, target name, title, author,
#  dir menu entry, description, category)
texinfo_documents = [
    (
        master_doc,
        "flyteidl",
        "flyteidl Documentation",
        author,
        "flyteidl",
        "Python SDK for Flyte (https://flyte.org).",
        "Miscellaneous",
    ),
]


# -- Extension configuration -------------------------------------------------
# intersphinx configuration
intersphinx_mapping = {
    "python": ("https://docs.python.org/{.major}".format(sys.version_info), None),
}


def setup(app):
    app.add_config_value(
        "recommonmark_config",
        {
            "auto_toc_tree_section": "Contents",
            "enable_math": False,
            "enable_inline_math": False,
            "enable_eval_rst": True,
        },
        True,
    )
    app.add_transform(AutoStructify)
